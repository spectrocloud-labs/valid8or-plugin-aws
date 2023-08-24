# valid8or-plugin-aws
The AWS [valid8or](https://github.com/spectrocloud-labs/valid8or) plugin ensures that your AWS environment matches a user-configurable expected state.

## Description
The AWS valid8or plugin reconciles `AwsValidator` custom resources to perform the following validations against your AWS environment:

1. Compare the IAM permissions associated with an IAM role against an expected permission set
2. Compare the usage for a particular service quota against the active quota
3. Compare the tags associated with a subnet against an expected tag set

Each `AwsValidator` CR is (re)-processed every two minutes to continuously ensure that your AWS environment matches the expected state.

See the [samples](https://github.com/spectrocloud-labs/valid8or-plugin-aws/tree/main/config/samples) directory for example `AwsValidator` configurations.

## Supported Service Quotas by AWS Service
EC2:
- EC2-VPC Elastic IPs

ELB:
- Application Load Balancers per Region
- Classic Load Balancers per Region

VPC:
- VPCs per Region
- Subnets per VPC
- NAT gateways per Availability Zone
- Network interfaces per Region
- Internet gateways per Region

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [kind](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=<some-registry>/valid8or-plugin-aws:tag
```

3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/valid8or-plugin-aws:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller from the cluster:

```sh
make undeploy
```

## Contributing
All contributions are welcome! Feel free to reach out on the [Spectro Cloud community Slack](https://spectrocloudcommunity.slack.com/join/shared_invite/zt-g8gfzrhf-cKavsGD_myOh30K24pImLA#/shared-invite/email).

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/),
which provide a reconcile function responsible for synchronizing resources until the desired state is reached on the cluster.

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

