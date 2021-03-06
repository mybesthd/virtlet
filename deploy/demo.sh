#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail
set -o errtrace

KUBE_VERSION="${KUBE_VERSION:-1.11}"
CRIPROXY_DEB_URL="${CRIPROXY_DEB_URL:-https://github.com/Mirantis/criproxy/releases/download/v0.12.0/criproxy-nodeps_0.12.0_amd64.deb}"
NONINTERACTIVE="${NONINTERACTIVE:-}"
NO_VM_CONSOLE="${NO_VM_CONSOLE:-}"
INJECT_LOCAL_IMAGE="${INJECT_LOCAL_IMAGE:-}"
dind_script="dind-cluster-v${KUBE_VERSION}.sh"
kubectl="${HOME}/.kubeadm-dind-cluster/kubectl"
BASE_LOCATION="${BASE_LOCATION:-https://raw.githubusercontent.com/Mirantis/virtlet/master/}"
RELEASE_LOCATION="${RELEASE_LOCATION:-https://github.com/Mirantis/virtlet/releases/download/}"
VIRTLET_DEMO_RELEASE="${VIRTLET_DEMO_RELEASE:-}"
VIRTLET_DEMO_BRANCH="${VIRTLET_DEMO_BRANCH:-}"
VIRTLET_ON_MASTER="${VIRTLET_ON_MASTER:-}"
VIRTLET_MULTI_NODE="${VIRTLET_MULTI_NODE:-}"
IMAGE_REGEXP_TRANSLATION="${IMAGE_REGEXP_TRANSLATION:-1}"
MULTI_CNI="${MULTI_CNI:-}"
# Convenience setting for local testing:
# BASE_LOCATION="${HOME}/work/kubernetes/src/github.com/Mirantis/virtlet"
cirros_key="demo-cirros-private-key"
# just initialize it
declare virtlet_release
declare virtlet_docker_tag

virtlet_nodes=()
if [[ ${VIRTLET_ON_MASTER} ]]; then
  virtlet_nodes+=(kube-master)
fi
if [[ !${VIRTLET_ON_MASTER} || ${VIRTLET_MULTI_NODE} ]]; then
  virtlet_nodes+=(kube-node-1)
fi
if [[ ${VIRTLET_MULTI_NODE} ]]; then
  virtlet_nodes+=(kube-node-2)
fi

# In case of linuxkit / moby linux, -v will not work so we can't
# mount /lib/modules and /boot.
using_linuxkit=
if ! docker info|grep -s '^Operating System: .*Docker for Windows' > /dev/null 2>&1 ; then
    if docker info|grep -s '^Kernel Version: .*-moby$' >/dev/null 2>&1 ||
         docker info|grep -s '^Kernel Version: .*-linuxkit-' > /dev/null 2>&1 ; then
        using_linuxkit=1
    fi
fi

function demo::step {
  local OPTS=""
  if [ "$1" = "-n" ]; then
    shift
    OPTS+="-n"
  fi
  GREEN="$1"
  shift
  if [ -t 2 ] ; then
    echo -e ${OPTS} "\x1B[97m* \x1B[92m${GREEN}\x1B[39m $*" >&2
  else
    echo ${OPTS} "* ${GREEN} $*" >&2
  fi
}

function demo::ask-before-continuing {
  if [[ ! ${NONINTERACTIVE} ]]; then
    echo "Press Enter to continue or Ctrl-C to stop." >&2
    read
  fi
}

function demo::ask-user {
  if [[ ${1:-} = "" ]]; then
    echo "no prompt message provided" >&2
    exit 1
  fi

  if [[ ${2:-} = "" ]]; then
    echo "no return var name provided" >&2
    exit 1
  fi

  local  __resultvar=$2
  local reply="false"
  while true; do
    read -p "$(tput bold)$(tput setaf 3) ${1} (yY/nN): $(tput sgr0)" reply
    case $reply in
      Y|y) reply="true"; break;;
      N|n) echo "Abort"; reply="false"; break;;
      *) echo "Please answer y[Y] or n[N].";;
    esac
  done
  eval $__resultvar="'$reply'"
}

function demo::get-dind-cluster {
  download="true"
  if [[ -f ${dind_script} ]]; then
    demo::step "Will update ${dind_script} script to the latest version"
    if [[ ! ${NONINTERACTIVE} ]]; then
        demo::ask-user "Do you want to redownload ${dind_script} ?" download
        if [[ ${download} = "true" ]]; then
          rm "${dind_script}"
        fi
    else
       demo::step "Will now clear existing ${dind_script}"
       rm "${dind_script}"
    fi
  fi

  if [[ ${download} = "true" ]]; then
    demo::step "Will download ${dind_script} into current directory"
    demo::ask-before-continuing
    wget "https://raw.githubusercontent.com/kubernetes-sigs/kubeadm-dind-cluster/master/fixed/${dind_script}"
    chmod +x "${dind_script}"
  fi
}

function demo::get-cirros-ssh-keys {
  if [[ -f ${cirros_key} ]]; then
    return 0
  fi
  demo::step "Will download ${cirros_key} into current directory"
  wget -O ${cirros_key} "https://raw.githubusercontent.com/Mirantis/virtlet/${virtlet_release}/examples/vmkey"
  chmod 600 ${cirros_key}
}

function demo::start-dind-cluster {
  demo::step "Will now clear any kubeadm-dind-cluster data on the current Docker"
  if [[ ! ${NONINTERACTIVE} ]]; then
    echo "Cirros ssh connection will be open after Virtlet setup is complete, press Ctrl-D to disconnect." >&2
  fi
  echo "To clean up the cluster, use './dind-cluster-v${KUBE_VERSION}.sh clean'" >&2
  demo::ask-before-continuing
  "./${dind_script}" clean
  "./${dind_script}" up
}

function demo::jq-patch {
  local node="${1}"
  local expr="${2}"
  local filename="${3}"
  docker exec "${node}" \
         bash -c "jq '${expr}' '${filename}' >/tmp/jqpatch.tmp && mv /tmp/jqpatch.tmp '${filename}'"
}

function demo::install-cni-genie {
  "${kubectl}" apply -f https://docs.projectcalico.org/v2.6/getting-started/kubernetes/installation/hosted/kubeadm/1.6/calico.yaml
  demo::wait-for "Calico etcd" demo::pods-ready k8s-app=calico-etcd
  demo::wait-for "Calico node" demo::pods-ready k8s-app=calico-node
  "${kubectl}" apply -f https://raw.githubusercontent.com/Huawei-PaaS/CNI-Genie/master/conf/1.8/genie-plugin.yaml

  demo::wait-for "CNI Genie" demo::pods-ready k8s-app=genie
  demo::jq-patch kube-node-1 '.cniVersion="0.3.0"|.default_plugin="calico,flannel"' /etc/cni/net.d/00-genie.conf
  demo::jq-patch kube-node-1 '.cniVersion="0.3.0"' /etc/cni/net.d/10-calico.conf
  demo::jq-patch kube-node-1 '.cniVersion="0.3.0"' /etc/cni/net.d/10-flannel.conflist
}

function demo::install-cri-proxy {
  local virtlet_node="${1}"
  demo::step "Installing CRI proxy package on ${virtlet_node} container"
  docker exec "${virtlet_node}" /bin/bash -c "curl -sSL '${CRIPROXY_DEB_URL}' >/criproxy.deb && dpkg -i /criproxy.deb && rm /criproxy.deb"
}

function demo::fix-mounts {
  local virtlet_node="${1}"
  demo::step "Marking mounts used by virtlet as shared in ${virtlet_node} container"
  docker exec "${virtlet_node}" mount --make-shared /dind
  docker exec "${virtlet_node}" mount --make-shared /dev
  if [[ ! ${using_linuxkit} ]]; then
    docker exec "${virtlet_node}" mount --make-shared /boot
  fi
  docker exec "${virtlet_node}" mount --make-shared /sys/fs/cgroup
}

function demo::inject-local-image {
  local virtlet_node="${1}"
  demo::step "Copying local mirantis/virtlet image into ${virtlet_node} container"
  docker save mirantis/virtlet | docker exec -i "${virtlet_node}" docker load
}

function demo::label-and-untaint-node {
  local virtlet_node="${1}"
  demo::step "Applying label to ${virtlet_node}:" "extraRuntime=virtlet"
  "${kubectl}" label node "${virtlet_node}" extraRuntime=virtlet
  if [[ ${VIRTLET_ON_MASTER} ]]; then
      demo::step "Checking/removing master taint from ${virtlet_node}"
    if [[ $("${kubectl}" get node kube-master -o jsonpath='{.spec.taints[?(@.key=="node-role.kubernetes.io/master")]}') ]]; then
      "${kubectl}" taint nodes kube-master node-role.kubernetes.io/master-
    fi
  fi
}

function demo::pods-ready {
  local label="$1"
  local out
  if ! out="$("${kubectl}" get pod -l "${label}" -n kube-system \
                           -o jsonpath='{ .items[*].status.conditions[?(@.type == "Ready")].status }' 2>/dev/null)"; then
    return 1
  fi
  if ! grep -v False <<<"${out}" | grep -q True; then
    return 1
  fi
  return 0
}

function demo::service-ready {
  local name="$1"
  if ! "${kubectl}" describe service -n kube-system "${name}"|grep -q '^Endpoints:.*[0-9]\.'; then
    return 1
  fi
}

function demo::wait-for {
  local title="$1"
  local action="$2"
  local what="$3"
  shift 3
  demo::step "Waiting for:" "${title}"
  while ! "${action}" "${what}" "$@"; do
    echo -n "." >&2
    sleep 1
  done
  echo "[done]" >&2
}

virtlet_pod=
function demo::virsh {
  local opts=
  if [[ ${1:-} = "console" ]]; then
    # using -it with `virsh list` causes it to use \r\n as line endings,
    # which makes it less useful
    local opts="-it"
  fi
  if [[ ! ${virtlet_pod} ]]; then
    virtlet_pod=$("${kubectl}" get pods -n kube-system -l runtime=virtlet -o name|head -1|sed 's@.*/@@')
  fi
  "${kubectl}" exec ${opts} -n kube-system "${virtlet_pod}" -c virtlet -- virsh "$@"
}

function demo::ssh {
  local cirros_ip=

  demo::get-cirros-ssh-keys

  if [[ ! ${virtlet_pod} ]]; then
    virtlet_pod=$("${kubectl}" get pods -n kube-system -l runtime=virtlet -o name|head -1|sed 's@.*/@@')
  fi

  if [[ ! ${cirros_ip} ]]; then
    while true; do
      cirros_ip=$("${kubectl}" get pod cirros-vm -o jsonpath="{.status.podIP}")
      if [[ ! ${cirros_ip} ]]; then
        echo "Waiting for cirros IP..."
        sleep 1
        continue
      fi
      echo "Cirros IP is ${cirros_ip}."
      break
    done
  fi

  echo "Trying to establish ssh connection to cirros-vm..."
  while ! internal::ssh ${virtlet_pod} ${cirros_ip} "echo Hello" | grep -q "Hello"; do
    sleep 1
    echo "Trying to establish ssh connection to cirros-vm..."
  done

  echo "Successfully established ssh connection. Press Ctrl-D to disconnect."
  internal::ssh ${virtlet_pod} ${cirros_ip}
}

function internal::ssh {
  virtlet_pod=${1}
  cirros_ip=${2}
  shift 2

  ssh -oProxyCommand="${kubectl} exec -i -n kube-system ${virtlet_pod} -c virtlet -- nc -q0 ${cirros_ip} 22" \
    -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -q \
    -i ${cirros_key} cirros@cirros-vm "$@"
}

function demo::vm-ready {
  local name="$1"
  # note that the following is not a bulletproof check
  if ! demo::virsh list --name | grep -q "${name}\$"; then
    return 1
  fi
}

function demo::kvm-ok {
  demo::step "Checking for KVM support..."
  # The check is done inside the node container because it has proper /lib/modules
  # from the docker host. Also, it'll have to use mirantis/virtlet image
  # later anyway.
  if [[ ${using_linuxkit} ]]; then
    return 1
  fi
  # use kube-master node as all of the DIND nodes in the cluster are similar
  if ! docker exec kube-master docker run --privileged --rm -v /lib/modules:/lib/modules "mirantis/virtlet:${virtlet_docker_tag}" kvm-ok; then
    return 1
  fi
}

function demo::get-correct-virtlet-release {
  # will use most recently published virtlet release
  # (virtlet releases are pre-releases now so not returned in /latest)
  local  __resultvar=$1
  local jq_filter=".[0].tag_name"
  local last_release
  last_release=$(curl --silent https://api.github.com/repos/Mirantis/virtlet/releases | docker exec -i kube-master jq "${jq_filter}" | sed 's/^\"\(.*\)\"$/\1/')
  if [[ $__resultvar ]]; then
    eval $__resultvar="'$last_release'"
  else
    echo "$last_release"
  fi
}

function demo::start-virtlet {
  local -a virtlet_config=(--from-literal=download_protocol=http --from-literal=image_regexp_translation="$IMAGE_REGEXP_TRANSLATION")
  if [[ ${VIRTLET_DEMO_BRANCH} ]]; then
    if [[ ${VIRTLET_DEMO_BRANCH} = "master" ]]; then
      virtlet_release="master"
      virtlet_docker_tag="latest"
    else
      virtlet_release="${VIRTLET_DEMO_BRANCH}"
      virtlet_docker_tag=$(echo $VIRTLET_DEMO_BRANCH | sed -e "s/\//_/g")
      BASE_LOCATION="https://raw.githubusercontent.com/Mirantis/virtlet/${virtlet_release}/"
    fi
  else
    if [[ ${VIRTLET_DEMO_RELEASE} ]]; then
      virtlet_release="${VIRTLET_DEMO_RELEASE}"
    else
      demo::get-correct-virtlet-release virtlet_release
    fi
    # set correct urls and names
    virtlet_docker_tag="${virtlet_release}"
    BASE_LOCATION="https://raw.githubusercontent.com/Mirantis/virtlet/${virtlet_release}/"
  fi
  echo "Will run demo using Virtlet:${virtlet_release} for demo and ${virtlet_docker_tag} as docker tag"
  if demo::kvm-ok; then
    demo::step "Setting up Virtlet configuration with KVM support"
  else
    demo::step "Setting up Virtlet configuration *without* KVM support"
    virtlet_config+=(--from-literal=disable_kvm=y)
  fi

  "${kubectl}" create configmap -n kube-system virtlet-config "${virtlet_config[@]}"
  # new functionality added post 0.8.2
  # that logic could be removed later
  if [[ ${BASE_LOCATION} == https://* ]]; then
    # remote location so fetch file
    rm -f demo_images.yaml
    status_code=$(curl -w "%{http_code}" --silent -o demo_images.yaml "${BASE_LOCATION}"/deploy/images.yaml)
    if [[ $status_code == "200" ]]; then
      "${kubectl}" create configmap -n kube-system virtlet-image-translations --from-file demo_images.yaml
    fi
  else
    if [[ -f ${BASE_LOCATION}/deploy/images.yaml ]]; then
      "${kubectl}" create configmap -n kube-system virtlet-image-translations --from-file "${BASE_LOCATION}/deploy/images.yaml"
    fi
  fi

  demo::step "Deploying Virtlet DaemonSet with docker tag ${virtlet_docker_tag}"
  docker run --rm "mirantis/virtlet:${virtlet_docker_tag}" virtletctl gen --tag "${virtlet_docker_tag}" |
      "${kubectl}" apply -f -
  demo::wait-for "Virtlet DaemonSet" demo::pods-ready runtime=virtlet
}

function demo::start-nginx {
  "${kubectl}" run nginx --image=nginx --expose --port 80
}

function demo::start-vm {
  demo::step "Starting sample CirrOS VM"
  "${kubectl}" create -f "${BASE_LOCATION}/examples/cirros-vm.yaml"
  demo::wait-for "CirrOS VM" demo::vm-ready cirros-vm
  if [[ ! "${NO_VM_CONSOLE:-}" ]]; then
    demo::step "Establishing ssh connection to the VM. Use Ctrl-D to disconnect"
    demo::ssh
  fi
}

if [[ ${1:-} = "--help" || ${1:-} = "-h" ]]; then
  cat <<EOF >&2
Usage: ./demo.sh

This script runs a simple demo of Virtlet[1] using kubeadm-dind-cluster[2]
ssh connection will be established after Virtlet setup is complete, Ctrl-D
can be used to disconnect from it.
Use 'curl http://nginx.default.svc.cluster.local' from VM console to test
cluster networking.

To clean up the cluster, use './dind-cluster-v${KUBE_VERSION}.sh clean'
[1] https://github.com/Mirantis/virtlet
[2] https://github.com/kubernetes-sigs/kubeadm-dind-cluster
EOF
  exit 0
fi

demo::get-dind-cluster
if [[ ${MULTI_CNI} ]]; then
  export NUM_NODES=1
  export CNI_PLUGIN=flannel
fi
demo::start-dind-cluster
if [[ ${MULTI_CNI} ]]; then
  demo::install-cni-genie
fi
for virtlet_node in "${virtlet_nodes[@]}"; do
  demo::fix-mounts "${virtlet_node}"
  demo::install-cri-proxy "${virtlet_node}"
  if [[ ${INJECT_LOCAL_IMAGE:-} ]]; then
    demo::inject-local-image "${virtlet_node}"
  fi
  demo::label-and-untaint-node "${virtlet_node}"
done
demo::start-virtlet
demo::start-nginx
demo::start-vm
