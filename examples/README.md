# Running the examples

## Create a Codespace (Optional)

Optionally, you can run the examples using a Dev Container in GitHub Codespaces. To do so, click the "+" at the root of the repo, then follow the steps to launch a new Codespace using the provided devcontainer config.

The default pre-configured devcontainer comes with all of the tools you will need already preinstalled, including: Terraform, Node/NPM, Prettier, TFLint, etc.

If you prefer not to use the pre-defined Devcontainer, you can use any suitable dev environment where Terraform is installed.

## Example Prereqs

1. A Google Sheet to use as the load Destination.
1. Your Google Creds so that Airbyte can authenticate to the Google Sheet.
1. Airbyte API Key to use when reading and writing Airbyte config using the Airbyte API.

### Running the Example

```console
echo "Switching to the example project directory..."
cd examples/pokeapi

echo "Initializing the Terraform provider..."
terraform init

echo "Printing the deployment plan..."
echo "You will be prompted for input variables as needed."
terraform plan

echo "Applying the plan and deploying to the Airbyte service..."
terraform apply
```
