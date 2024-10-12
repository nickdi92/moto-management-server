"use client"

import {useRouter} from "next/navigation";
import {GetUserDataFromLocalStorage, IsUserLoggedIn} from "@/app/helpers/userHelper";
import DashboardHeader from "@/app/components/dashboard/layout";
import React, {useState} from "react";
import {UpdateUser} from "@/app/api/apiUsers";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import MotoLoader from "@/app/components/loader";


export default function User() {
    const router = useRouter();
    if (!IsUserLoggedIn()) {
        router.push("/admin/login");
    }
    const [isLoading, setIsLoading] = useState(false);
    const userInfo = GetUserDataFromLocalStorage();
    /**
     * Modifier for each property of user
     */
    const [name, setName] = useState(userInfo.name ?? "");
    const [lastname, setLastname] = useState(userInfo.lastname ?? "");
    const [street, setStreet] = useState(userInfo.address.street) ?? "";
    const [city, setCity] = useState(userInfo.address.city) ?? "";
    const [province, setProvince] = useState(userInfo.address.province ?? "");
    const [state, setState] = useState(userInfo.address.state ?? "");
    const [zipCode, setZipCode] = useState(userInfo.address.zip_code ?? "");
    const [fiscalCode, setFiscalCode] = useState(userInfo.registry.fiscal_code ?? "");
    const [dob, setDob] = useState(userInfo.registry.dob ?? "")
    
    const handleInputChange = (action, value) => {
        action(value);
    }
    
    async function onSubmit(event) {
        event.preventDefault();
        setIsLoading(true);
        try {
            const formData = new FormData(event.target);
            let bodyRaw = {};
            for (const pair of formData.entries()) {
                if (pair[0].includes(".")) {
                    let keys = pair[0].split(".");
                    if (!bodyRaw.hasOwnProperty(keys[0])) {
                        bodyRaw[keys[0]] = {};
                    }
                    bodyRaw[keys[0]][keys[1]] = pair[1];
                } else {
                    bodyRaw[pair[0]] = pair[1];
                }
            }
            bodyRaw["email"] = userInfo.email;
            bodyRaw["username"] = userInfo.username;
            bodyRaw["password"] = userInfo.password;
            
            const user = await UpdateUser(bodyRaw);
            
            if (user.status_code === 200) {
                toast.success("User updated successfully", {
                    autoClose: 1000
                });
            } else {
                console.error("USER updated", user);
                toast.error("User update error. Receive error code: " + user.status_code, {
                    autoClose: 1000
                });
            }
        } catch (error) {
            toast.error("[userProfile] - Error: " + error, {
                autoClose: 2000
            });
            console.error("[userProfile] - Error: ", error);
        }
        
        setIsLoading(false);
    }
    
    return (
        <>
            <div className="min-h-full bg-white shadow">
                {<DashboardHeader title={name + " " + lastname}/>}
                
                <main className="bg-white shadow min-h-screen">
                    <div className="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8 ">
                        <form onSubmit={onSubmit} className="relative">
                            <div className="space-y-12">
                                <div className="border-b border-gray-900/10 pb-12">
                                    <h2 className="text-base font-semibold leading-7 text-gray-900">
                                        Personal Information
                                    </h2>
                                    <div className="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-9 mb-10">
                                        <div className="sm:col-span-3">
                                            <label htmlFor="name"
                                                   className="block text-sm font-medium leading-6 text-gray-900">
                                                First name
                                            </label>
                                            <div className="mt-2">
                                                <input
                                                    id="name"
                                                    name="name"
                                                    value={name}
                                                    onChange={(e) => handleInputChange(setName, e.target.value)}
                                                    type="text"
                                                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                />
                                            </div>
                                        </div>
                                        
                                        <div className="sm:col-span-3">
                                            <label htmlFor="lastname"
                                                   className="block text-sm font-medium leading-6 text-gray-900">
                                                Last name
                                            </label>
                                            <div className="mt-2">
                                                <input
                                                    id="lastname"
                                                    name="lastname"
                                                    value={lastname}
                                                    onChange={(e) => handleInputChange(setLastname, e.target.value)}
                                                    type="text"
                                                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                />
                                            </div>
                                        </div>
                                        
                                        <div className="sm:col-span-3">
                                            <label htmlFor="email"
                                                   className="block text-sm font-medium leading-6 text-gray-900">
                                                Email address
                                            </label>
                                            <div className="mt-2">
                                                <input
                                                    disabled={true}
                                                    id="email"
                                                    name="email"
                                                    type="email"
                                                    value={userInfo.email}
                                                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                />
                                            </div>
                                        </div>
                                    </div>
                                    <h2 className="text-base font-semibold leading-7 text-gray-900">
                                        Address Information
                                    </h2>
                                    <div className="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-9 mb-10">
                                        <div className="sm:col-span-3">
                                            <label htmlFor="street"
                                                   className="block text-sm font-medium leading-6 text-gray-900">
                                                Street address
                                            </label>
                                            <div className="mt-2">
                                                <input
                                                    id="street"
                                                    name="address.street"
                                                    value={street}
                                                    onChange={(e) => handleInputChange(setStreet, e.target.value)}
                                                    type="text"
                                                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                />
                                            </div>
                                        </div>
                                        
                                        <div className="sm:col-span-3">
                                            <label htmlFor="city"
                                                   className="block text-sm font-medium leading-6 text-gray-900">
                                                City
                                            </label>
                                            <div className="mt-2">
                                                <input
                                                    id="city"
                                                    name="address.city"
                                                    type="text"
                                                    value={city}
                                                    onChange={(e) => handleInputChange(setCity, e.target.value)}
                                                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                />
                                            </div>
                                        </div>
                                        
                                        <div className="sm:col-span-3">
                                            <label htmlFor="province"
                                                   className="block text-sm font-medium leading-6 text-gray-900">
                                                Province
                                            </label>
                                            <div className="mt-2">
                                                <input
                                                    id="province"
                                                    name="address.province"
                                                    value={province}
                                                    onChange={(e) => handleInputChange(setProvince, e.target.value)}
                                                    type="text"
                                                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                />
                                            </div>
                                        </div>
                                        
                                        <div className="sm:col-span-3">
                                            <label htmlFor="zip_code"
                                                   className="block text-sm font-medium leading-6 text-gray-900">
                                                ZIP / Postal code
                                            </label>
                                            <div className="mt-2">
                                                <input
                                                    id="zip_code"
                                                    name="address.zip_code"
                                                    value={zipCode}
                                                    onChange={(e) => handleInputChange(setZipCode, e.target.value)}
                                                    type="text"
                                                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                />
                                            </div>
                                        </div>
                                        <div className="sm:col-span-3">
                                            <label htmlFor="state"
                                                   className="block text-sm font-medium leading-6 text-gray-900">
                                                Country
                                            </label>
                                            <div className="mt-2">
                                                <input
                                                    id="state"
                                                    name="address.state"
                                                    value={state}
                                                    onChange={(e) => handleInputChange(setState, e.target.value)}
                                                    type="text"
                                                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                />
                                            </div>
                                        </div>
                                    </div>
                                    
                                    <h2 className="text-base font-semibold leading-7 text-gray-900">
                                        Registry Information
                                    </h2>
                                    <div className="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-9 mb-10">
                                        <div className="sm:col-span-3">
                                            <label htmlFor="fiscal_code"
                                                   className="block text-sm font-medium leading-6 text-gray-900">
                                                Fiscal code
                                            </label>
                                            <div className="mt-2">
                                                <input
                                                    id="fiscal_code"
                                                    name="registry.fiscal_code"
                                                    value={fiscalCode}
                                                    onChange={(e) => handleInputChange(setFiscalCode, e.target.value)}
                                                    type="text"
                                                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                />
                                            </div>
                                        </div>
                                        
                                        <div className="sm:col-span-3">
                                            <label htmlFor="dob"
                                                   className="block text-sm font-medium leading-6 text-gray-900">
                                                Date Of Birth
                                            </label>
                                            <div className="mt-2">
                                                <input
                                                    id="dob"
                                                    name="registry.dob"
                                                    type="text"
                                                    value={dob}
                                                    onChange={(e) => handleInputChange(setDob, e.target.value)}
                                                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                />
                                            </div>
                                        </div>
                                    </div>
                                
                                </div>
                            </div>
                            <div className="mt-6 flex items-center justify-end gap-x-6">
                                <button type="button" className="text-sm font-semibold leading-6 text-gray-900">
                                    Cancel
                                </button>
                                <button
                                    type="submit"
                                    className="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                                >
                                    Save
                                </button>
                            </div>
                        </form>
                    </div>
                </main>
            </div>
            <ToastContainer />
        </>
    )
}