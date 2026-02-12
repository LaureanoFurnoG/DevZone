import Keycloak from "keycloak-js";

const keycloak = new Keycloak({
    url: "http://localhost:8081",
    realm: "devzone",
    clientId: "devzone-frontend",
})

export default keycloak