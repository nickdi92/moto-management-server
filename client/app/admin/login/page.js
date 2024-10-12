'use client'

import React, {useState} from "react";
import { useRouter } from 'next/navigation';
import {LoginUser} from "@/app/api/apiUsers";
import {GetUserFullName} from "@/app/helpers/userHelper";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import Logo from "@/app/components/logo";
import MotoLoader from "@/app/components/loader";

export default function Login() {
    const [isLoading, setIsLoading] = useState(false);
    const router = useRouter()
    
    async function onSubmit(event) {
        event.preventDefault();
        setIsLoading(true);
        try {
            const formData = new FormData(event.target);
            let bodyRaw = {};
            for (const pair of formData.entries()) {
                bodyRaw[pair[0]] = pair[1];
            }
            const loginUser = await LoginUser(bodyRaw);
            
            if (loginUser.hasOwnProperty("user") && loginUser.user) {
                toast.success("Login avvenuto con successo. Benvenuto " + GetUserFullName(), {
                    autoClose: 1500
                });
                setTimeout(function () {
                    router.push("/admin/dashboard");
                }, 1500)
            } else {
                toast.error("Error on login User. Status code: " + loginUser.status_code)
                console.error("Error on login User. Status code: ", loginUser.status_code);
            }
        } catch (error) {
            toast.error("Error on login User: " + error)
            console.error("Error on login User: ", error);
        } finally {
            setIsLoading(false);
        }
        event.target.reset();
    }
    
    return (
        <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
            <div className="sm:mx-auto sm:w-full sm:max-w-sm">
                {<Logo context="guest" />}
                <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
                    Sign in to your account
                </h2>
            </div>
            <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
                <form action="#" method="POST" className="space-y-6 relative" onSubmit={onSubmit}>
                    {isLoading ? <MotoLoader /> : ""}
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
            <ToastContainer />
        </div>
    )
}
  