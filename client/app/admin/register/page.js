'use client'

import React, {useState} from "react";
import {setCookie} from "cookies-next";
import NotificationSuccess from "@/app/components/notifications/success";
import NotificationError from "@/app/components/notifications/error";
import Logo from "@/app/components/logo";
import {CreateUser} from "@/app/api/apiUsers";

export default function Register() {
    const [isLoading, setIsLoading] = useState(false);
    const [userCreated, setUserCreated] = useState(false);
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
            const data = CreateUser(bodyRaw);
            
            if (data.token) {
                setUserCreated(true);
                setUserCreationHasErrors(false);
                setNotificationMessage("Registrazione avvenuta con sucesso !");
            } else {
                console.error("Error on creating new User: ", data.registerRouteErr);
                setUserCreationHasErrors(true);
                setNotificationMessage("Error on creating new User: " + data.registerRouteErr);
            }
        } catch (error) {
            console.error("Error on creating new User: ", error);
            setUserCreationHasErrors(true);
            setNotificationMessage("Error on creating new User: " + error)
        } finally {
            setIsLoading(false);
        }
        event.target.reset();
    }
    
    return (
        <>
            
            <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
                <div className="sm:mx-auto sm:w-full sm:max-w-sm">
                    {<Logo context="guest"/>}
                    <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
                        Create New Account
                    </h2>
                </div>
                
                <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
                    {userCreated ? <NotificationSuccess message={notificationMessage}/> : ""}
                    {userCreationHasErrors ? <NotificationError message={notificationMessage}/> : ""}
                    <form action="#" method="POST" className="space-y-6" onSubmit={onSubmit}>
                        <div>
                            <label htmlFor="name" className="block text-sm font-medium leading-6 text-gray-900">
                                Name
                            </label>
                            <div className="mt-2">
                                <input
                                    id="name"
                                    name="name"
                                    type="text"
                                    required
                                    autoComplete="name"
                                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                />
                            </div>
                        </div>
                        
                        <div>
                            <label htmlFor="lastname" className="block text-sm font-medium leading-6 text-gray-900">
                                Lastname
                            </label>
                            <div className="mt-2">
                                <input
                                    id="lastname"
                                    name="lastname"
                                    type="text"
                                    required
                                    autoComplete="lastname"
                                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                />
                            </div>
                        </div>
                        
                        <div>
                            <label htmlFor="email" className="block text-sm font-medium leading-6 text-gray-900">
                                Email address
                            </label>
                            <div className="mt-2">
                                <input
                                    id="email"
                                    name="email"
                                    type="email"
                                    required
                                    autoComplete="email"
                                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                />
                            </div>
                        </div>
                        
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
                                {isLoading ? "Creazione utente in corso ..." : "Crea account"}
                            </button>
                            <a href="/admin/login"
                               className="flex w-full justify-center rounded-md bg-green-600 mt-4 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Log
                                In</a>
                        </div>
                    </form>
                </div>
            </div>
        </>
    )
}