# url-shortener
This is a basic URL shortener done as a learning exercise on the following topics:
- Programming with Go
- Working with Firestore for the first time
- More work on CI/CD and Infrastructure as Code concepts
- More work on deploying APIs to Google Cloud Platform.

A frontend for this is still to come, but I did this project for a [Boot.dev](https://boot.dev) submission, and they are focused on backend work.

#### Running locally
To run locally:
- Clone the repo
- Install the [Google Cloud CLI](https://cloud.google.com/sdk/gcloud)
- Run `go build` from the `application` directory
- Run `gcloud emulators firestore start --host-port=127.0.0.1:<PORT>`
- Run your built executable.

A `.env` file is provided - typically these are gitignored, but this doesn't contain any sensitive info. Make sure the `PORT` environment variable matches the port you chose when running the Firestore emulator.

#### Deploying to the cloud
You could deploy all of this manually in the following fashion:
- Set up a Google Cloud Platform project
- Run the [OpenTofu](https://opentofu.org) project's dev environment from `infrastructure/environments/dev`
  - `tofu init`
  - `tofu plan`
  - `tofu apply`
- Run `go build` from the `application` directory
- Deploy to Cloud Run using `gcloud run deploy` from the `application` directory.

CI/CD pipelines are provided, so feel free to fork this repo and set it up. Run the `dev` Tofu project first to turn on all the required APIs. This project uses a single GCP project with two environments, as per Google's [GitOps documentation](https://cloud.google.com/docs/terraform/resource-management/managing-infrastructure-as-code#setting_up_your_github_repository) and [sample repo](https://github.com/GoogleCloudPlatform/solutions-terraform-cloudbuild-gitops). They don't control service APIs with Terraform, which is what makes this part complicated for this repo.

The appendix tables below define all the variables required to run OpenTofu, manual deploys to Cloud Run, and the secrets required if using CI/CD.

#### Appendix A: OpenTofu variables
If deploying manually, save these to `terraform.tfvars` in the `infrastructure/environments/dev` directory.
| Name | Purpose |
|:--   |:--      |
| `project_id` | Name of your GCP project |
| `env` | "dev" when running manually, "prod" for CI/CD |
| `app_name` | Whatever you want to call the app. I used `url-shortener` |
| `sa_provisioner_name` | Service account used for all Terraform provisioning |
| `github_repo_id` | Your Github repo internal ID if deploying via CI/CD with workload identity federation |
| `github_repo_owner_id` | Your Github account internal ID if deploying via CI/CD with workload identity federation |

#### Appendix B: Manual deployments to Cloud Run
If deploying manually, the following arguments will be required:
| Name | Purpose |
|:--   |:--      |
| `PROJECT_ID` | Your GCP project name |
| `DATABASE_ID` | The name of your database in Firestore |
| `DOMAIN_NAME` | The URL of your Cloud Run service |

The value for `DOMAIN_NAME` will need to be set using `gcloud run services update` once deployed for the first time and you have the URL. Have a look at how the CI/CD pipeline does this for details.

#### Appendix C: CI/CD
When deploying via CI/CD, use the following secrets in your repo:
| Name | Purpose |
|:--   |:--      |
| `CLOUD_RUN_SA` | The service account to run the service with. Must have the correct permissions |
| `FIRESTORE_DB` | The name of your target Firestore database. Just the name, not the fully qualified /projects/firestore/database etc. |
| `GCP_PROJECT` | Your GCP project name - required for deployment target |
| `GCP_PROJECT_NUMBER` | Your GCP project number - required for Cloud Run deployment namespace |
| `GCP_REGION` | Your target GCP region - note, the OpenTofu code currently hardcodes `australia-southeast1` |
| `GCP_SA_NAME` | Service account being used to deploy to Cloud Run. Different from the one running the service! |
| `GCP_STATE_BUCKET` | Target Cloud Storage bucket to store Tofu state in - recommend creating manually |
| `GCP_WORKLOAD_PROVIDER` | The name of your GCP Workload Identity Provider created by Terraform |
| `REPO_ID` | Same as `github_repo_id` in Appendix A |
| `REPO_OWNER_ID` | Same as `github_repo_owner_id` in Appendix A |

#### Appendix D: GCP permissions
I recommend the following IAM roles for the different service accounts to get things running:

##### Cloud Run SA
- Cloud Datastore User
- Cloud Run Invoker

##### GCP SA for deployment
- Artifact Registry Administrator
- Cloud Datastore Owner
- Cloud Run Admin
- Service Account User
- Storage Admin

Service Account Token Creator may also be required.