## Policy API

### PolicyException status

[kyverno-policy-operator](https://github.com/giantswarm/kyverno-policy-operator) writes the `status` of a `PolicyException` (`gspolex`):

- `observedGeneration`: the `metadata.generation` the status describes.
- `conditions`:
  - `Ready`: whether the exception covers what was asked for: every Kyverno PolicyException generated for it is applied and every target is translated. Reasons: `Reconciled`, `InvalidNamespace`, `NameTaken`, `LookupFailed`, `ApplyFailed`, `DeleteFailed`, `UnsupportedKind`.
    When several checks fail, `Ready` takes the first failing reason in that order, and its message lists all of them.
  - `PoliciesResolved`: whether every listed policy matches a CEL policy. It is informational and does not affect `Ready`, so health checks on `Ready` (Flux `wait`, kstatus) are not blocked by policies still being migrated. Reasons:
    - `Resolved`: every listed policy matches a CEL policy.
    - `NotMigrated`: a listed policy matches only a legacy ClusterPolicy. The legacy exception covers it, and the CEL exception takes over once the policy is migrated. This is expected during the migration.
    - `PolicyNotFound`: a listed policy matches no policy at all, for example a typo, a removed policy, or a policy not installed yet. It wins over `NotMigrated`, and the message names the policies for each.
  - `TargetsTranslated`: whether every target can be expressed in a CEL exception. Reasons: `Translated`, `UnsupportedKind`.
- `generatedExceptions`: the `apiVersion`, `namespace` and `name` of each Kyverno PolicyException generated for it.
- `unresolvedPolicies`: listed policies that match no CEL policy, whether they match only a legacy ClusterPolicy or no policy at all.
- `unsupportedTargetKinds`: target kinds left out of the CEL exception, such as `Pod/exec`.

`kubectl get gspolex` shows the `Ready` condition's status and reason.
