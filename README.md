# booking-app

A Go project built to explore core language concepts including goroutines, WaitGroups, maps, structs, slices, and package management.

---

## Go Modules, Packages, and Dependency Management — Notes

### 1. Packages in Go

* Every .go file must declare a package.
* Packages are **directory-based**, not file-based.
* All files in the same directory must have the same package name.
* A package represents a logical unit of code.

---

### 2. Module (go.mod)

* A module is the root of a Go project.
* Defined by a go.mod file.
* Example:

  `go
  module github.com/username/project-name
  `
* The module path must match the repository path if the code is intended to be imported.

---

### 3. Import Paths

* Import paths are derived from:

  `
  module path + folder path
  `
* Example:

  `go
  import "github.com/username/project/auth"
  `
* Go does not import individual files, only packages (folders).

---

### 4. internal Directory

* internal/ is a special directory.
* Packages inside internal/:

  * Can only be imported within the same module.
  * Cannot be accessed from outside the module.
* Used to enforce encapsulation.

---

### 5. go get

* Used to add dependencies to a module.
* Must be executed **inside a directory that has go.mod**.
* Fetches modules or packages, not individual files.

---

### 6. go install

* Used to install binaries globally.
* Not used for adding dependencies to a project.

---

### 7. go.sum

* Stores checksums of dependencies.
* Ensures integrity and reproducibility.
* Automatically managed by Go; not edited manually.

---

### 8. go mod tidy

`go mod tidy` analyzes your code and your `go.mod` file. It adds any missing dependencies you're using but didn't explicitly list, and it removes any dependencies you no longer need. In short, it ensures your `go.mod` and `go.sum` are consistent with the actual imports in your project.

* Adds missing dependencies that are imported in your code but not yet listed in `go.mod`.
* Removes unused dependencies that are listed in `go.mod` but no longer imported.
* Keeps `go.mod` and `go.sum` consistent with the actual imports in your project.

---

### 9. Publishing a Go Package

To make a package usable by others:

1. Initialize module:

   `ash
   go mod init github.com/username/project
   `
2. Ensure correct module path in go.mod.
3. Push code to a public repository.
4. (Optional but recommended) Tag versions.

---

### 10. Common Issues

#### Module Path Mismatch

* If go.mod contains:

  `go
  module projectname
  `

  but imported as:

  `
  github.com/username/project
  `

  → Go will fail.

#### Private or Non-existent Repository

* Repository must exist and be accessible.

#### Using internal Packages Externally

* Not allowed by design.

#### Missing go.mod

* Without it, Go cannot treat the project as a module.

---

### 11. Key Rules Summary

* Module path must match repository path.
* Packages are folders, not files.
* internal packages are private.
* go get works only inside a module.
* go.mod defines identity; go.sum ensures integrity.

---

### 12. Fix — Module Path Mismatch & Go Proxy Cache Issue

When this repo originally had module bookingapp in go.mod, importing it from another project failed:

`
module declares its path as: bookingapp
        but was required as: github.com/sharafath-ali/booking-app
`

**Root cause:** The module name in go.mod did not match the GitHub repository path.

**Fix applied:**

1. Updated go.mod:

   `go
   module github.com/sharafath-ali/booking-app
   `

2. Updated internal import in main.go:

   `go
   import "github.com/sharafath-ali/booking-app/helper"
   `

3. Ran go mod tidy and pushed.

**Secondary issue — Go Proxy Cache:**

Even after the fix, the Go module proxy (proxy.golang.org) had cached the old go.mod. To bypass the cache and fetch the corrected version directly from GitHub:

`powershell
="direct"
="*"
go get github.com/sharafath-ali/booking-app/helper@c558906
`

* GOPROXY=direct — skips the public proxy and fetches straight from GitHub.
* GONOSUMDB=* — skips checksum database verification (needed for direct non-tagged fetches).
* @c558906 — pins to the exact commit that contains the corrected go.mod.

> **Tip:** Tag your releases (git tag v0.1.0 && git push --tags) to avoid pseudo-version issues. Tagged versions are handled more reliably by the Go proxy.

