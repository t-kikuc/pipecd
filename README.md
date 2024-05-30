[![Build](https://github.com/pipe-cd/pipecd/actions/workflows/build.yaml/badge.svg)](https://github.com/pipe-cd/pipecd/actions/workflows/build.yaml)
[![Test](https://github.com/pipe-cd/pipecd/actions/workflows/test.yaml/badge.svg)](https://github.com/pipe-cd/pipecd/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/pipe-cd/pipecd)](https://goreportcard.com/report/github.com/pipe-cd/pipecd)
[![codecov](https://codecov.io/gh/pipe-cd/pipecd/branch/master/graph/badge.svg)](https://codecov.io/gh/pipe-cd/pipecd)
[![Release](https://img.shields.io/github/v/release/pipe-cd/pipecd?label=Release)](https://github.com/pipe-cd/pipecd/releases/latest)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](/LICENSE)
[![Documentation](https://img.shields.io/badge/Documentation-pipecd-informational.svg)](https://pipecd.dev/docs/)
[![Slack](https://img.shields.io/badge/Slack-%23pipecd-informational.svg)](https://app.slack.com/client/T08PSQ7BQ/C01B27F9T0X)
[![Twitter](https://img.shields.io/twitter/url/https/twitter.com/pipecd_dev.svg?style=social&label=Follow%20%40pipecd_dev)](https://twitter.com/pipecd_dev)

<p align="center">
  <img src="https://github.com/pipe-cd/pipecd/blob/master/docs/static/images/logo.png" width="180"/>
</p>

<p align="center">
  A GitOps style continuous delivery platform that provides consistent deployment and operations experience for any applications
  <br/>
  <a href="https://pipecd.dev"><strong>Explore PipeCD docs »</strong></a>
  <a href="https://play.pipecd.dev?project=play"><strong>Play with live demo »</strong></a>
</p>

#

![](https://github.com/pipe-cd/pipecd/blob/master/docs/static/images/rolled-back-deployment.png)

### Overview

PipeCD provides __a unified continuous delivery solution for multiple application kinds on multi-cloud__ that empowers engineers to deploy faster with more confidence, a GitOps tool that enables doing deployment operations by pull request on Git.

![](https://github.com/pipe-cd/pipecd/blob/master/docs/static/images/pipecd-explanation.png)

#

### Why PipeCD?

- Simple, unified and easy to use but powerful pipeline definition to construct your deployment
- Same deployment interface to deploy applications of any platform, including Kubernetes, Terraform, GCP Cloud Run, AWS Lambda, AWS ECS
- No CRD or applications' manifest changes are required; Only need a pipeline definition along with your application manifests
- No deployment credentials are exposed or required outside the application cluster
- Built-in deployment analysis as part of the deployment pipeline to measure impact based on metrics, logs, emitted requests
- Easy to interact with any CI; The CI tests and builds artifacts, PipeCD takes the rest
- Insights show metrics like lead time, deployment frequency, MTTR and change failure rate to measure delivery performance
- Designed to manage thousands of cross-platform applications in multi-cloud for company scale but also work well for small projects

And many more, please explore in [docs](https://pipecd.dev/docs).

#

### Quickstart

The [quickstart guide](https://pipecd.dev/docs/quickstart/) shows how to set up PipeCD components and deploy a hello-world application with PipeCD.

#

### Installation

The [installation guide](https://pipecd.dev/docs/installation/) explains and helps set up PipeCD for your real-life production environment.

#

### Release and development

Please check [the release page](https://github.com/pipe-cd/pipecd/releases) to see what is included in the latest release. Also, [Releases documentation](https://github.com/pipe-cd/pipecd/blob/master/RELEASES.md/) explains our versioning and release cycle.

Go to the [Discussion](https://github.com/pipe-cd/pipecd/discussions) to know what we are working on and which will be added to the next release.

#

### Contributing

We'd love you to join us! Please see the [Contributing docs](CONTRIBUTING.md).

#

### Community

As a CNCF Sandbox project, PipeCD follows [CNCF Code of Conduct](https://github.com/pipe-cd/pipecd/blob/master/CODE_OF_CONDUCT.md).

#

### Adopters

You can find a list of publicly recognized users of the PipeCD in the [ADOPTERS.md](ADOPTERS.md) file. We strongly encourage all PipeCD users to add their names to this list, as we love to see the community's growing success!

#

### Thanks to the contributors of PipeCD!


#

**We are a [Cloud Native Computing Foundation](https://cncf.io/) sandbox project.**

<img src="https://www.cncf.io/wp-content/uploads/2022/07/cncf-color-bg.svg" width=300 />

The Linux Foundation® (TLF) has registered trademarks and uses trademarks. For a list of TLF trademarks, see [Trademark Usage](https://www.linuxfoundation.org/trademark-usage/).
