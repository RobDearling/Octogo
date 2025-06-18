<p align="center">
  <img alt="logo" src="./docs/images/logo.png" height=250>
</p>

---

![GitHub Release (latest SemVer)](https://img.shields.io/github/v/release/RobDearling/octogo?logo=github&label=Release&sort=semver)

Octogo is a Golang package that provides a client for interfacing with the Octopus Energy API.

## Getting Started with Your API Key

### Prerequisites

Before you can use Octogo, you'll need an Octopus Energy account and an API key. All API endpoints from this package require authentication via an API key.

### Generating Your API Key

1. **Log into your Octopus Energy account** at [octopus.energy](https://octopus.energy)
2. **Navigate to your dashboard** and go to Account Settings
3. **Access API settings** by visiting the [API Access page](https://octopus.energy/dashboard/new/accounts/personal-details/api-access)
4. **Generate your key** by clicking "Generate API Key" or similar button
5. **Copy and store your key securely** - you'll need this for authentication

## Examples

### Electricity Meter

```go
client := octogo.NewClient("my-secure-api-key")
ctx := context.TODO()
meter, resp, err := client.Meter.GetElectricityMeter(ctx, mpan)

if err != nil {
  fmt.Printf("An error occurred while fetching the electricity meter: %v\n", err)
  return err
}


if resp.StatusCode == http.StatusOK {
  fmt.Printf("Meter Details:\n")
  fmt.Printf("  GSP: %s\n", meter.GSP)
  fmt.Printf("  MPAN: %s\n", meter.MPAN)
  fmt.Printf("  Profile Class: %d\n", meter.ProfileClass)
}
```
