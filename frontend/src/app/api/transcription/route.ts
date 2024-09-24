import { NextResponse } from 'next/server';

const BACKEND_URL = process.env.NEXT_PUBLIC_BACKEND_URL; // GolangのバックエンドURLを環境変数から取得

export async function POST(request: Request) {
    const formData = await request.formData();
    const jobName = formData.get('jobName');
    const file = formData.get('file') as File;  // ファイルを取得
    const languageCode = formData.get('languageCode');
    const isSpeakerSeparationEnabled = formData.get('isSpeakerSeparationEnabled') === 'true';
    const isTimestampEnabled = formData.get('isTimestampEnabled') === 'true';

    try {
        // 1. フロントエンドからGolangバックエンドにファイルを直接送信
        const uploadFormData = new FormData();
        uploadFormData.append('file', file);

        const uploadResponse = await fetch(`${BACKEND_URL}/api/s3/upload`, {
            method: 'POST',
            body: uploadFormData,
        });

        if (!uploadResponse.ok) {
            throw new Error('Failed to upload file to the backend');
        }

        // アップロードが成功したか確認
        const uploadData = await uploadResponse.json();
        const s3FileUrl = uploadData.url;

        // 2. S3にアップロードされたファイルを使ってTranscribeジョブを開始
        const transcribeResponse = await fetch(`${BACKEND_URL}/api/transcriptions/start`, {  // Transcribeジョブ開始用のエンドポイント
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                jobName: jobName,
                mediaUri: s3FileUrl,  // アップロードされたファイルのURLを送信
                languageCode: languageCode,
                // customVocabularyName: "doda",
                // isSpeakerSeparationEnabled: isSpeakerSeparationEnabled,
                // isTimestampEnabled: isTimestampEnabled,
            }),
        });

        // Transcribe jobが重複している場合、409ステータスコードを処理
        if (transcribeResponse.status === 409) {
            return NextResponse.json({
                message: 'Job name already exists',
            }, { status: 409 });
        }

        if (!transcribeResponse.ok) {
            throw new Error('Failed to start transcription job on the backend');
        }

        const transcribeData = await transcribeResponse.json();

        return NextResponse.json({
            message: 'Transcription job started successfully',
            ...transcribeData,
        }, { status: 200 });

    } catch (error) {
        // console.error('Error in transcription process:', error);

        // エラーメッセージを返す
        if (error instanceof Error) {
            return NextResponse.json({
                message: 'An error occurred during transcription process.',
                error: error.message
            }, { status: 500 });
        } else {
            return NextResponse.json({
                message: 'An unknown error occurred during transcription process.'
            }, { status: 500 });
        }
    }
}
