import {GetAuthenticationHeader, UpdateBearerToken, UpdateBearerTokenExpiration} from "@/app/api/auth";
import {IsUserLoggedIn, SetIsUserLoggedIn, UpdateUserDataToLocalStorage, SetUsername} from "@/app/helpers/userHelper";

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
        SetUsername(bodyRaw["username"]);
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

export async function GetUserInfo(bodyRaw) {
    const call = await fetch("http://localhost:8080/admin/user/get", {
        method: "POST",
        body: JSON.stringify(bodyRaw),
        headers: {
            "Content-type": "application/json",
            "Authorization": GetAuthenticationHeader()
        }
    });
    
    const userInfo = await call.json();
    if (userInfo.status_code === 200 && userInfo.user) {
        UpdateUserDataToLocalStorage(userInfo.user);
    }
    
    return userInfo;
}