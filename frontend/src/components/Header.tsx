import React from 'react';

const Header: React.FC = () => {
    return (
        <header className="bg-gradient-to-r from-blue-500 to-purple-500 shadow-md">
            <div className="container mx-auto px-4 py-4 flex items-center justify-between">
                <h1 className="text-2xl font-bold text-white">My Application</h1>
                <nav className="hidden md:flex space-x-6">
                    <a href="#" className="text-white hover:text-gray-200 transition duration-200">ホーム</a>
                    <a href="#" className="text-white hover:text-gray-200 transition duration-200">サービス</a>
                    <a href="#" className="text-white hover:text-gray-200 transition duration-200">お問い合わせ</a>
                </nav>
                <div className="flex items-center space-x-4">
                    <button className="bg-white text-gray-800 px-3 py-1 rounded-md shadow-sm hover:bg-gray-100 transition duration-200">ログイン</button>
                    <button className="text-white hover:text-gray-200 transition duration-200">
                        <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth="2">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M12 4v16m8-8H4" />
                        </svg>
                    </button>
                </div>
            </div>
        </header>
    );
};

export default Header;
