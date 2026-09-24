import pykube
import pytest
from pytest_helm_charts.clusters import Cluster

crd_names = [
    "automatedexceptions.policy.giantswarm.io",
    "policies.policy.giantswarm.io",
    "policyconfigs.policy.giantswarm.io",
    "policyexceptions.policy.giantswarm.io",
    "policymanifests.policy.giantswarm.io",
]


def get_crd(kube_cluster: Cluster, name: str) -> pykube.CustomResourceDefinition:
    return pykube.CustomResourceDefinition.objects(kube_cluster.kube_client).get_by_name(name)


@pytest.mark.smoke
def test_api_working(kube_cluster: Cluster) -> None:
    """Test that we can connect to the Kubernetes API."""
    assert kube_cluster.kube_client is not None
    assert len(pykube.Node.objects(kube_cluster.kube_client)) >= 1


@pytest.mark.smoke
@pytest.mark.parametrize("name", crd_names)
def test_crd_exists(kube_cluster: Cluster, name: str) -> None:
    """Test that the chart installed the CRD."""
    assert get_crd(kube_cluster, name).obj["spec"]["group"] == "policy.giantswarm.io"


@pytest.mark.smoke
def test_policyexception_status(kube_cluster: Cluster) -> None:
    """Test that PolicyException has the status subresource and its printer columns."""
    crd = get_crd(kube_cluster, "policyexceptions.policy.giantswarm.io")
    version = next(v for v in crd.obj["spec"]["versions"] if v["name"] == "v1alpha1")

    assert "status" in version.get("subresources", {})
    columns = {c["name"] for c in version.get("additionalPrinterColumns", [])}
    assert {"Ready", "Reason", "Policies"} <= columns
