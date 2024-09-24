'use client';

import React, { useState } from 'react';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { Badge } from '@/components/ui/badge';
import { Globe, CheckCircle, Clock } from 'lucide-react';
import TranscriptionJobStatus from "@/components/TranscriptionJobStatus";
import { languages } from "@/utils/languages";  // 言語リストを取得

interface TranscriptionJob {
    jobName: string;
    creationTime: string;
    completionTime: string | null;
    languageCode: string;
    transcriptionJobStatus: string;
}

interface Props {
    jobs: TranscriptionJob[];
}

const TranscriptionList: React.FC<Props> = ({ jobs }) => {
    const [selectedJob, setSelectedJob] = useState<string | null>(null);

    // 言語コードから言語名を取得する関数
    const getLanguageName = (code: string) => {
        const language = languages.find((lang) => lang.value === code);
        return language ? language.label : 'その他';
    };

    return (
        <div className="max-w-5xl mx-auto py-10">
            <Card className="shadow-lg border border-gray-200 rounded-lg">
                <CardHeader className="bg-gradient-to-r from-purple-500 to-indigo-600 text-white p-4 rounded-t-lg mb-5">
                    <CardTitle className="text-2xl font-semibold">ジョブ一覧</CardTitle>
                </CardHeader>
                <CardContent className="">
                    <Table className="w-full border-collapse table-auto bg-white shadow-sm">
                        <TableHeader>
                            <TableRow className="bg-gray-50 text-left text-gray-600 border-b border-gray-200">
                                <TableHead className="py-3 px-4">ジョブ名</TableHead>
                                <TableHead className="py-3 px-4">作成日時</TableHead>
                                <TableHead className="py-3 px-4">言語 <Globe className="inline h-4 w-4 text-gray-500" /></TableHead>
                                <TableHead className="py-3 px-4">ステータス</TableHead>
                            </TableRow>
                        </TableHeader>
                        <TableBody>
                            {jobs.map((job) => (
                                <TableRow
                                    key={job.jobName}
                                    onClick={() => setSelectedJob(job.jobName)}
                                    className="cursor-pointer hover:bg-gray-100 transition-colors border-b border-gray-200"
                                >
                                    <TableCell className="py-3 px-4 text-gray-900">{job.jobName}</TableCell>
                                    <TableCell className="py-3 px-4 text-gray-600">{job.creationTime}</TableCell>
                                    <TableCell className="py-3 px-4 text-gray-600">{getLanguageName(job.languageCode)}</TableCell>
                                    <TableCell className="py-3 px-4">
                                        <TranscriptionJobStatus status={job.transcriptionJobStatus} />
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>

                    {selectedJob && (
                        <div className="mt-6">
                            <Card className="bg-blue-50 border border-blue-200">
                                <CardContent>
                                    <p className="text-blue-700">選択されたジョブ: <span className="font-bold">{selectedJob}</span></p>
                                </CardContent>
                            </Card>
                        </div>
                    )}
                </CardContent>
            </Card>
        </div>
    );
};

export default TranscriptionList;
