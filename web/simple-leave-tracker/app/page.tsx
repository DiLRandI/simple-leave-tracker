export default function Home() {
  return (
    <>
        <section id='hero' className='bg-green-700 py-20 text-white'>
            <div className='container mx-auto flex items-center px-4'>
                <div className='w-1/2'>
                    <h1 className='mb-6 text-6xl font-bold'>Simplify Your <span className='text-yellow-300'>HR Tasks</span></h1>
                    <p className='mb-8 text-xl'>Efficient HR solutions for modern businesses. Focus on what matters most — your
                        people.
                    </p>
                    <div className='space-y-4'><a href='#services'
                        className='rounded-full bg-yellow-400 px-6 py-3 font-semibold text-green-900 border border-transparent mr-2'>Explore
                        Our Services</a><a href='#contact'
                        className='rounded-full border-2 border-white bg-transparent px-6 py-3 font-semibold text-white'>Get in
                        Touch</a>
                    </div>
                </div>
                <div className='relative w-1/2'><img src='https://placehold.jp/3d4070/ffffff/600x400.png' alt='HR solutions'
                        className='rounded-lg shadow-2xl'></img>
                </div>
            </div>
        </section>
        <section id='about' className='bg-gray-100 py-20'>
        <div className='container mx-auto px-4 text-center'>
            <h2 className='mb-4 text-4xl font-bold text-green-900'>About Us</h2>
            <p className='mx-auto mb-12 max-w-2xl text-xl text-gray-600'>At SimpleHR, we believe in making HR processes simple,
            fast, and efficient. Our team of experts is committed to delivering top-notch HR solutions tailored to your
            business needs.</p>
        </div>
        </section>
    </>
  );
}
