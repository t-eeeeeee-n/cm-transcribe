import React from 'react';
import Client from './client';

const Page: React.FC = () => {
    // ここで必要なデータを取得やサーバーサイド処理を行うことができます
    // 例えば、デフォルトの言語やその他の設定などを取得する処理を行うことができます。

    return (
        <div>
            {/* サーバーコンポーネントからクライアントコンポーネントを呼び出す */}
            <Client />
        </div>
    );
};

export default Page;