mod "local" {
  color = "#f19428ff"
  require {
    plugin "kubernetes" {
      min_version = "1.5.0"
    }
    mod "github.com/turbot/steampipe-mod-kubernetes-insights" {
      version = "*"
    }
  }
}