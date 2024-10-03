'use client'

import React, {useState, Suspense} from "react";
import { setCookie, getCookie } from "cookies-next";
import NotificationSuccess from "@/app/components/notifications/success";
import NotificationError from "@/app/components/notifications/error";

export default function Login() {
    const [isLoading, setIsLoading] = useState(false);
    const [userLoggedInd, setUserIsLoggedIn] = useState(false);
    const [notificationMessage, setNotificationMessage] = useState("");
    const [userCreationHasErrors, setUserCreationHasErrors] = useState(false);

    async function onSubmit(event) {
        event.preventDefault();
        setIsLoading(true);
        try {
          const formData = new FormData(event.target);
          let bodyRaw = {};
          for (const pair of formData.entries()) {
            bodyRaw[pair[0]] = pair[1];
          }
          if (!getCookie("bearer_token")) {
              const cookie = await fetch("http://localhost:8080/admin/user/refresh-token", {
                  method: "POST",
                  body: JSON.stringify(bodyRaw),
                  headers: {
                      "Content-type": "application/json",
                  }
              });
              const data = await cookie.json();
              if (data.token) {
                  setCookie("bearer_token", data.token);
                  setCookie("bearer_token_expiration", data.expire_at);
              } else {
                  console.error("Error on refreshing User token.");
                  setUserCreationHasErrors(true);
                  setNotificationMessage("Error on refreshing User token.")
              }
          }

          const response = await fetch("http://localhost:8080/admin/user/login", {
              method: "POST",
              body: JSON.stringify(bodyRaw),
              headers: {
                "Content-type": "application/json",
                "Authorization": "Bearer " + getCookie("bearer_token")
              }
          });


          const data = await response.json();
          if (data.token) {
            setCookie("bearer_token", data.token);
            setCookie("bearer_token_expiration", data.expire_at);
            setUserIsLoggedIn(true);
            setUserCreationHasErrors(false);
            setNotificationMessage("Benvenuto " + bodyRaw["username"])
          } else {
            console.error("Error on login User: ", data.loginRouteErr);
            setUserCreationHasErrors(true);
            setNotificationMessage("Error on login User: " + data.loginRouteErr)
          }
        } catch (error) {
          console.error("Error on login User: ", error);
          setUserCreationHasErrors(true);
          setNotificationMessage("Error on login User: " + error)
        } finally {
          setIsLoading(false);
        }
        event.target.reset();
    }

    return (
          <Suspense fallback={<h1>Loading feed...</h1>}>
            <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
              <div className="sm:mx-auto sm:w-full sm:max-w-sm">
                <img
                  alt="Your Company"
                  src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
                  className="mx-auto h-10 w-auto"
                />
                <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
                  Sign in to your account
                </h2>
              </div>
              <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
                { userLoggedInd ? <NotificationSuccess  message={notificationMessage} /> : "" }
                { userCreationHasErrors ? <NotificationError message={notificationMessage} /> : "" }
                <form action="#" method="POST" className="space-y-6" onSubmit={onSubmit}>
                  <div>
                    <label htmlFor="username" className="block text-sm font-medium leading-6 text-gray-900">
                      Username
                    </label>
                    <div className="mt-2">
                      <input
                        id="username"
                        name="username"
                        type="text"
                        required
                        autoComplete="username"
                        className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                      />
                    </div>
                  </div>

                  <div>
                    <div className="flex items-center justify-between">
                      <label htmlFor="password" className="block text-sm font-medium leading-6 text-gray-900">
                        Password
                      </label>
                      <div className="text-sm">
                        <a href="#" className="font-semibold text-indigo-600 hover:text-indigo-500">
                          Forgot password?
                        </a>
                      </div>
                    </div>
                    <div className="mt-2">
                      <input
                        id="password"
                        name="password"
                        type="password"
                        required
                        autoComplete="current-password"
                        className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                      />
                    </div>
                  </div>

                  <div>
                    <button
                      type="submit"
                      disabled={isLoading}
                      className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                    >
                      Sign in
                    </button>
                    <a href="/admin/register"
                       className="flex w-full justify-center rounded-md bg-green-600 mt-4 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
                        Create an account
                    </a>
                  </div>
                </form>
              </div>
            </div>
          </Suspense>
    )
  }
  