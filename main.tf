provider "solace" {
    host = "localhost:8080"
    base_path = "/SEMP/v2/config"
    user = "admin"
    password = "alex"
    msg_vpn = "go-tf-1"
}

resource "solace_msgvpn" "my-vpn" {
    name = "go-tf-1"
    # enabled = true
    # authentication_basic_enabled = false
    max_spool_usage = 1150
}

resource "solace_aclprofile" "my-acl" {
    name = "ach-test-acl-1"
    msg_vpn = "go-tf-1"
    client_connection_default_action = "disallow"
    publish_topic_default_action = "allow"
    subscribe_topic_default_action = "allow"
}
