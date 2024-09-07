import type { Metadata } from "next";
import localFont from "next/font/local";
import "./globals.css";

const geistSans = localFont({
  src: "./fonts/GeistVF.woff",
  variable: "--font-geist-sans",
  weight: "100 900",
});
const geistMono = localFont({
  src: "./fonts/GeistMonoVF.woff",
  variable: "--font-geist-mono",
  weight: "100 900",
});

export const metadata: Metadata = {
  title: "Simple leave tracker",
  description: "Simple leave tracker created by DiLRandI",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        <header className='bg-white shadow-md'>
    <div className='container mx-auto flex items-center justify-between px-4 py-6'>
      <div className='text-3xl font-bold text-green-600'>SimpleHR</div>
      <nav className='space-x-6'><a href='#about' className='text-gray-600'>About</a><a href='#services'
          className='text-gray-600'>Services</a><a href='#contact' className='text-gray-600'>Contact</a></nav><button
        className='rounded-full bg-green-600 px-4 py-2 text-white'>Login</button>
    </div>
  </header>

        {children}

        <footer className='bg-gray-800 py-8 text-white'>
    <div className='container mx-auto flex items-center justify-between px-4'>
      <div className='mb-4 text-2xl font-bold'>SimpleHR</div>
      <nav className='space-x-4'><a href='#' className='hover:text-green-400'>Privacy Policy</a><a href='#'
          className='hover:text-green-400'>Terms of Service</a><a href='#' className='hover:text-green-400'>Contact Us</a></nav>
    </div>
    <div className='mt-8 text-center text-gray-400'>
      <p>© 2024 SimpleHR. All rights reserved.</p>
    </div>
  </footer>
      </body>

    </html>





  );
}
