import React, { useState } from 'react';
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion";
import UploadCSV from "@/components/UploadCSV";
import VocabularyTable from "@/components/VocabularyTable";
import {Vocabulary} from "@/types/Vocabulary";

interface DictionarySettingsProps {
    vocabularies: Vocabulary[];
    onVocabularyChange: (index: number, field: keyof Vocabulary, value: string) => void;
    onAddVocabulary: () => void;
    onRemoveVocabulary: (index: number) => void;
    setVocabulary: (vocabularies: Vocabulary[]) => void;
}

const DictionarySettings: React.FC<DictionarySettingsProps> = ({
                                                                   vocabularies,
                                                                   onVocabularyChange,
                                                                   onAddVocabulary,
                                                                   onRemoveVocabulary,
                                                                   setVocabulary,
                                                               }) => {
    const [isDirectInputOpen, setDirectInputOpen] = useState<boolean>(true);
    const [isCsvImportOpen, setCsvImportOpen] = useState<boolean>(true);
    const [fileName, setFileName] = useState<string | null>(null);

    const handleUpload = (uploadedVocabularies: Vocabulary[]) => {
        setVocabulary(uploadedVocabularies);
    };

    return (
        <Card className="shadow-lg border border-gray-200 rounded-lg bg-white dark:bg-gray-800 transition-shadow duration-300 ease-in-out hover:shadow-xl">
            <CardHeader className="p-4 bg-gradient-to-r from-purple-500 to-indigo-600 text-white rounded-t-lg">
                <CardTitle className="text-2xl font-semibold">辞書の設定</CardTitle>
            </CardHeader>
            <CardContent className="p-6">
                {/* 直接入力アコーディオン */}
                <Accordion type="multiple" defaultValue={["direct-input", "csv-import"]}>
                    <AccordionItem value="direct-input">
                        <AccordionTrigger onClick={() => setDirectInputOpen(!isDirectInputOpen)}>
                            直接入力
                        </AccordionTrigger>
                        <AccordionContent>
                            {isDirectInputOpen && (
                                <VocabularyTable
                                    vocabularies={vocabularies}
                                    onVocabularyChange={onVocabularyChange}
                                    onAddVocabulary={onAddVocabulary}
                                    onRemoveVocabulary={onRemoveVocabulary}
                                />
                            )}
                        </AccordionContent>
                    </AccordionItem>

                    {/* CSVインポートアコーディオン */}
                    <AccordionItem value="csv-import">
                        <AccordionTrigger onClick={() => setCsvImportOpen(!isCsvImportOpen)}>
                            CSVインポート
                        </AccordionTrigger>
                        <AccordionContent>
                            {isCsvImportOpen && (
                                <UploadCSV
                                    onUpload={handleUpload}
                                    fileName={fileName}
                                    setFileName={setFileName}
                                />
                            )}
                        </AccordionContent>
                    </AccordionItem>
                </Accordion>
            </CardContent>
        </Card>
    );
};

export default DictionarySettings;
