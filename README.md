## Policy API

### PolicyException status

[kyverno-policy-operator](https://github.com/giantswarm/kyverno-policy-operator) writes the `status` of a `PolicyException` (`gspolex`):

- `observedGeneration`: the `metadata.generation` the status describes.
- `conditions`:
  - `Ready`: whether the exception covers what was asked for: every Kyverno PolicyException generated for it is applied and every target is translated. Reasons: `Reconciled`, `InvalidNamespace`, `NameTaken`, `LookupFailed`, `ApplyFailed`, `DeleteFailed`, `UnsupportedKind`.
    When several checks fail, `Ready` takes the first failing reason in that order, and its message lists all of them.
  - `PoliciesResolved`: whether every listed policy matches a CEL policy. Reasons: `Resolved`, `PolicyNotFound`.
  - `TargetsTranslated`: whether every target can be expressed in a CEL exception. Reasons: `Translated`, `UnsupportedKind`.
- `generatedExceptions`: the `apiVersion`, `namespace` and `name` of each Kyverno PolicyException generated for it.
- `unresolvedPolicies`: listed policies that match no CEL policy yet.
- `unsupportedTargetKinds`: target kinds left out of the CEL exception, such as `Pod/exec`.

`kubectl get gspolex` shows the `Ready` condition's status and reason.
