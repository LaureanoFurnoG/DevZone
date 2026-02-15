import { useEffect, useState } from "react";
import keycloak from "./keycloak";

import { AuthContext, type Me } from "./AuthContext";

export const AuthProvider = ({children}: {children: React.ReactNode}) => {
    const [isLoading, setLoading] = useState(true)
    const [isAuthenticated, setAuthenticated] = useState(false)
    const [token, setToken] = useState<string | undefined>(); 
    const [me, setMe] = useState<Me | null>(null);

    useEffect(() =>{
        keycloak.init({
            onLoad: "check-sso",
            pkceMethod: "S256",
            checkLoginIframe: false,
            silentCheckSsoRedirectUri:
            window.location.origin + "/silent-check-sso.html",
        })
        .then((authenticated) =>{
            setAuthenticated(authenticated)
            setToken(keycloak.token)
            if (keycloak.tokenParsed){
                const name = keycloak.tokenParsed.name 
                const lastname = keycloak.tokenParsed.lastname
                const email = keycloak.tokenParsed.email
                const token = keycloak.token
                const sub = keycloak.subject
                setMe({name, email, lastname, token, sub})
            }
            setLoading(false)
        }).catch(() =>{
            setLoading(false)
        })

        const interval = setInterval(() =>{
            if(keycloak.authenticated){
                keycloak.updateToken(60)
                .then((refreshed) =>{
                    if(refreshed){
                        setToken(keycloak.token)
                    }
                }).catch(() =>{
                    keycloak.login()
                })
            }
        }, 100000) //refresh each 10s
        return () => clearInterval(interval); //clean interval when disassemble the provider
    }, [])

    const login = () => keycloak.login()

    const logout = () => keycloak.logout({
        redirectUri: window.location.origin
    })

    return( 
        <AuthContext.Provider value={{isAuthenticated, isLoading, me, token, login, logout}}>
            {children}
        </AuthContext.Provider>
    )
}