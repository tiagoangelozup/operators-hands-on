# operators-hands-on

## Set up your environment

Create your KinD cluster:
```shell
kind create cluster --name operators-hands-on
```

## Project creation

Use the OperatorSDK to create a new operator project:
```shell
mkdir -p $HOME/dev/projects/hello-kubernetes-operator
cd $HOME/dev/projects/hello-kubernetes-operator
# obs: if your version is newer than 1.16, you can skip de validation --skip-go-version-check
operator-sdk init --domain zup.com --repo github.com/example/hello-kubernetes-operator
```

Create a new API and Controller
```shell
operator-sdk create api --group install --version v1alpha1 --kind HelloKubernetes --resource --controller
```

## API definition

Define your API that is represented in `api/v1alpha1/hellokubernetes_types.go` by `HelloKubernetesSpec` type:
```go
// HelloKubernetesSpec defines the desired state of HelloKubernetes
type HelloKubernetesSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Enabled bool `json:"enabled"`
}
```

Then update the generated code and the CRD manifests by running:
```shell
make generate manifests
```

Now you can install your new API by running:
```shell
make install
```

## Implement the Controller

The controller template file on method `Reconcile` at `controllers/hellokubernetes_controller.go`

## Reconcilers return options

The following are a few possible return options for a Reconciler:

-   With the error:
    ```go
    return ctrl.Result{}, err
    ```

-   Without an error:
    ```go
    return ctrl.Result{Requeue: true}, nil
    ```

-   Therefore, to stop the Reconcile, use:

    ```go
    return ctrl.Result{}, nil
    ```

-   Reconcile again after X time:

    ```go
     return ctrl.Result{RequeueAfter: nextRun.Sub(r.Now())}, nil
    ```