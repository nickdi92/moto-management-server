import Image from "next/image";
import GithubLogo from "./assets/images/github.png";

export default async function Home() {
  return (
    <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
      <main className="flex flex-col gap-8 row-start-2 items-center sm:items-start">
        <h1 className="text-xl text-center">Moto Management Server</h1>
        <div className="flex gap-4 items-center flex-col sm:flex-row text-center">
          <a href="/admin/register">Create Account</a>
          <a href="/admin/login" className="bg-indigo-500 px-5 py-2 inline-block text-white rounded-md">Log In</a>
        </div>
      </main>
      <footer className="row-start-3 flex gap-6 flex-wrap items-center justify-center">
        <a
          className="flex items-center gap-2 hover:underline hover:underline-offset-4"
          href="https://github.com/nickdi92"
          target="_blank"
          rel="noopener noreferrer"
        >
          <Image
            src={GithubLogo}
            alt="File icon"
            width={100}
            height={25}
          />
          Nickdi
        </a>
        
      </footer>
    </div>
  );
}
