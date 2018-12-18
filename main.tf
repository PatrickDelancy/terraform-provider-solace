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

resource "solace_msgvpn" "my-vpn-2" {
    name = "go-tf-2"
    enabled = true
    max_spool_usage = 980
}

resource "solace_msgvpn" "my-vpn-3-conflict" {
    name = "go-tf-1"
}