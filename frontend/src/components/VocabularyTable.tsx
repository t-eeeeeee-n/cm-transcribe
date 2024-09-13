import React from 'react';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Trash2 } from "lucide-react";
import { Vocabulary } from "@/types/Vocabulary";

interface VocabularyTableProps {
    vocabularies: Vocabulary[];
    onVocabularyChange: (index: number, field: keyof Vocabulary, value: string) => void;
    onAddVocabulary: () => void;
    onRemoveVocabulary: (index: number) => void;
}

export default function VocabularyTable({
                                            vocabularies,
                                            onVocabularyChange,
                                            onAddVocabulary,
                                            onRemoveVocabulary,
                                        }: VocabularyTableProps) {
    return (
        <div className="mt-6 space-y-6">
            <div className="rounded-lg border border-gray-200 shadow-md overflow-hidden">
                <Table className="min-w-full divide-y divide-gray-200">
                    <TableHeader className="bg-gray-50">
                        <TableRow>
                            <TableHead className="w-[200px] py-4 px-4 text-left text-sm font-semibold text-gray-700">フレーズ（必須）</TableHead>
                            <TableHead className="w-[200px] py-4 px-4 text-left text-sm font-semibold text-gray-700">みたいに聞こえる（任意）</TableHead>
                            <TableHead className="w-[200px] py-4 px-4 text-left text-sm font-semibold text-gray-700">IPA（任意）</TableHead>
                            <TableHead className="w-[200px] py-4 px-4 text-left text-sm font-semibold text-gray-700">表示フレーズ（任意）</TableHead>
                            <TableHead className="w-[50px] py-4 px-4 text-left text-sm font-semibold text-gray-700"></TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody className="bg-white divide-y divide-gray-100">
                        {vocabularies.length === 0 ? (
                            <TableRow>
                                <TableCell colSpan={5} className="py-4 px-4 text-center text-gray-500">
                                    <div className="min-h-[60px] flex flex-col justify-center"> {/* Add fixed height for consistency */}
                                        <p>行はまだ追加されていません</p>
                                        <p>直接入力するか、CSVインポートで作成</p>
                                    </div>
                                </TableCell>
                            </TableRow>
                        ) : (
                            vocabularies.map((row, index) => (
                                <TableRow key={index} className="hover:bg-gray-50 transition-colors duration-200">
                                    <TableCell className="py-4 px-4">
                                        <Input
                                            value={row.phrase}
                                            onChange={(e) => onVocabularyChange(index, 'phrase', e.target.value)}
                                            required
                                            className="border-gray-300 focus:border-blue-500 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
                                        />
                                    </TableCell>
                                    <TableCell className="py-4 px-4">
                                        <Input
                                            value={row.soundsLike}
                                            onChange={(e) => onVocabularyChange(index, 'soundsLike', e.target.value)}
                                            className="border-gray-300 focus:border-blue-500 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
                                        />
                                    </TableCell>
                                    <TableCell className="py-4 px-4">
                                        <Input
                                            value={row.ipa}
                                            onChange={(e) => onVocabularyChange(index, 'ipa', e.target.value)}
                                            className="border-gray-300 focus:border-blue-500 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
                                        />
                                    </TableCell>
                                    <TableCell className="py-4 px-4">
                                        <Input
                                            value={row.displayAs}
                                            onChange={(e) => onVocabularyChange(index, 'displayAs', e.target.value)}
                                            className="border-gray-300 focus:border-blue-500 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
                                        />
                                    </TableCell>
                                    <TableCell className="py-4 px-4">
                                        <Button
                                            variant="ghost"
                                            size="icon"
                                            onClick={() => onRemoveVocabulary(index)}
                                            aria-label="行を削除"
                                            className="hover:text-red-600 hover:bg-red-50 transition-colors duration-200"
                                        >
                                            <Trash2 className="h-5 w-5" />
                                        </Button>
                                    </TableCell>
                                </TableRow>
                            ))
                        )}
                    </TableBody>
                </Table>
            </div>
            <div className="flex justify-end">
                <Button onClick={onAddVocabulary} className="bg-blue-600 text-white hover:bg-blue-700 transition-all duration-300">
                    行を追加
                </Button>
            </div>
        </div>
    );
}
