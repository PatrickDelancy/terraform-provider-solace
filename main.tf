provider "solace" {
    host = "localhost:8080"
    base_path = "/SEMP/v2/config"
    user = "admin"
    password = "alex"
}

resource "solace_msgvpn" "my-vpn" {
    name = "go-tf-1"
    enabled = false
    max_spool_usage = 1150
}
