import {GetAuthenticationHeader, UpdateBearerToken, UpdateBearerTokenExpiration} from "@/app/api/auth";
import {IsTokenExpired} from "@/app/api/token";

export async function CreateUser(bodyRaw) {
    const response = await fetch("http://localhost:8080/admin/user/create", {
        method: "POST",
        body: JSON.stringify(bodyRaw),
        headers: {
            "Content-type": "application/json"
        }
    });
    
    const data = await response.json();
    if (data.hasOwnProperty("token")) {
        UpdateBearerToken(data.token);
        UpdateBearerTokenExpiration(data.expire_at)
    }
    return data;
}

export async function LoginUser(bodyRaw) {
    if (!IsUserLoggedIn()) {
        const refreshToken = await RefreshUserToken(bodyRaw);
        if (refreshToken.hasOwnProperty("token")) {
            UpdateBearerToken(refreshToken.token);
            UpdateBearerTokenExpiration(refreshToken.expire_at)
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

    const loginResponse = await response.json()
    if (loginResponse.hasOwnProperty("user") && loginResponse.user) {
        SetIsUserLoggedIn(loginResponse.user)
    }
    return loginResponse;
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
    let userData = localStorage.getItem("user_data");
    let isUserLoggedIn = Boolean(localStorage.getItem("is_user_logged_in"));
    
    return userData && !IsTokenExpired(userData.token) && isUserLoggedIn;
}

export function SetIsUserLoggedIn(userData) {
    localStorage.setItem("user_data", JSON.stringify(userData));
    localStorage.setItem("is_user_logged_in", userData?.is_logged_in);
    UpdateBearerToken(userData.token);
}