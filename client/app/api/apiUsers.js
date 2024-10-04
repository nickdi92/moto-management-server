import {getCookie} from "cookies-next";
import {GetAuthenticationHeader, UpdateBearerToken} from "@/app/api/auth";

export async function CreateUser() {

}

export async function LoginUser(bodyRaw) {
    if (!IsUserLoggedIn()) {
        const refreshToken = RefreshUserToken(bodyRaw);
        if (refreshToken.hasOwnProperty("token")) {
            UpdateBearerToken(refreshToken.token);
        }
    }
    const response = await fetch("http://localhost:8080/admin/user/login", {
        method: "POST",
        body: JSON.stringify(bodyRaw),
        headers: {
            "Content-type": "application/json",
            "Authorization": GetAuthenticationHeader()
        }
    });

    return await response.json()
}

export async function RefreshUserToken(bodyRaw) {
    const cookie = await fetch("http://localhost:8080/admin/user/refresh-token", {
        method: "POST",
        body: JSON.stringify(bodyRaw),
        headers: {
            "Content-type": "application/json",
        }
    });
    return await cookie.json();
}

export function IsUserLoggedIn() {

}