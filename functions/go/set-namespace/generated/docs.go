// Code generated by "mdtogo"; DO NOT EDIT.
package generated

var SetNamespaceShort = `The ` + "`" + `set-namespace` + "`" + ` function replaces the ` + "`" + `namespace` + "`" + ` specific resource type in a variety
of KRM resources.`
var SetNamespaceLong = `
## Usage

This function replaces the KRM resources existing namespace to a new value.

### Target KRM resources

- This function updates all namespace-scoped KRM resources ` + "`" + `metadata.namespace` + "`" + ` fields. 
  We determine whether a KRM resource is namespace scoped by checking if it has ` + "`" + `metadata.namespace` + "`" + ` set and matches the "oldNamespace"
  If not, this function won't add new namespace.
- This function updates ` + "`" + `RoleBinding` + "`" + ` and ` + "`" + `ClusterRoleBinding` + "`" + ` resources ` + "`" + `subjects` + "`" + ` element whose kind is ` + "`" + `ServiceAccount` + "`" + `
  and the subject's ` + "`" + `namespace` + "`" + ` is set and matches the "oldNamespace".
- This function updates ` + "`" + `CustomResourceDefinition` + "`" + ` (CRD) ` + "`" + `spec/conversion/webhook/clientConfig/service/namespace` + "`" + ` field 
  if the field is set and matches the "oldNamespace"
- This function updates ` + "`" + `APIService` + "`" + ` ` + "`" + `spec/service/namespace` + "`" + ` field if the field is set and matches the "oldNamespace"
- This function updates the KRM resources annotation ` + "`" + `config.kubernetes.io/depends-on` + "`" + ` if this annotation contains the 
  matching namespace.

### FunctionConfig

This function supports the default ` + "`" + `ConfigMap` + "`" + ` as function config and a custom ` + "`" + `SetNamespace` + "`" + `. See below examples

ConfigMap as functionConfig
  apiVersion: v1
  kind: ConfigMap
  data:
    namespace: newNamespace # required
    namespaceMatcher: example # update namespace whose value is "example" to "newNamespace"

SetNamespace as functionConfig
  apiVersion: fn.kpt.dev/v1alpha1
  kind: SetNamespace
  namespace: newNamespace # required
  namespaceMatcher: example # update namespace whose value is "example" to "newNamespace"


### Three updating modes

This function supports three modes to flexibly choose and update the target namespaces.

# Restrict Mode:
All target KRM resources namespace has to have the same value. All namespace will be updated to the new value. 

ConfigMap as functionConfig
  apiVersion: v1
  kind: ConfigMap
  data:
    namespace: newNamespace # update all namespace fields to "newNamespace"

# DefaultNamespace Mode:

The input ` + "`" + `resourcelist.items` + "`" + ` contains one and only one namespace object. The function matches the namespace ` + "`" + `metadata.name` + "`" + `
with all other KRM resources, and only update the namespace if it matches the namespace object. 
If more than one namespace objects are found, raise errors;

  kind: ResourceList
  functionConfig:
    apiVersion: v1
    kind: ConfigMap
    data:
      namespace: newNs
  items:
  - apiVersion: v1
    kind: Namespace
    metadata:
      name: example # updated to "newNs"
  - apiVersion: v1
    kind: Service
    metadata:
      name: the-service1
      namespace: example  # updated to "newNs"
  - apiVersion: v1
    kind: Service
    metadata:
      name: the-service2
      namespace: irrelevant # skip since namespace does not match "example".

# Selector Mode:

Only updates the namespace which matches a given value. The "oldNamespace" refers to the argument in FunctionConfig

ConfigMap as functionConfig
  apiVersion: v1
  kind: ConfigMap
  data:
    namespace: newNamespace
    namespaceMatcher: example # update namespace whose value is "example" to "newNamespace"

SetNamespace as functionConfig
  apiVersion: fn.kpt.dev/v1alpha1
  kind: SetNamespace
  namespace: newNamespace
  namespaceMatcher: example # update namespace whose value is "example" to "newNamespace"

### DependsOn annotation

DependsOn annotation is a [kpt feature](https://kpt.dev/reference/annotations/depends-on/). This function updates the 
namespace segment in a depends-on annotation if matches oldNs
`
