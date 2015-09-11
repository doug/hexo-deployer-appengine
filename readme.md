#hexo deployment to Google AppEngine

You need to have google cloud sdk already installed. 

Download from: https://cloud.google.com/sdk/

(Make sure you've included the python sdk - if not, you can get it from here: https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Python)

Or install: `curl https://cloud.google.com/sdk/ | bash`

Also, you need to have setup your appengine project and authorized gcloud.

Setup your project on appengine at: https://console.developers.google.com/project

Authorize gcloud with: `gcloud auth login`

Add the new deploy to your `_deploy.yaml`

```yaml
## Docs: http://hexo.io/docs/deployment.html
deploy:
  type: appengine
  project: required_project_name
  version: optional_version_name
  password: optional_password
  # dryrun: true # useful for testing appengine locally
```
