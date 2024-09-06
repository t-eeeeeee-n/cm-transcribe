import React, { useState } from 'react';
import {
    Box,
    Button,
    Paper,
    Typography,
    Tabs,
    Tab,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    TextField,
    IconButton,
    Snackbar,
    Alert,
    Card,
    CardContent,
} from '@mui/material';
import DeleteIcon from '@mui/icons-material/Delete';
import UploadFileIcon from '@mui/icons-material/UploadFile';
import CloseIcon from '@mui/icons-material/Close';
import { read, utils } from 'xlsx';

interface Phrase {
    phrase: string;
    soundsLike: string;
    ipa: string;
    displayAs: string;
}

interface PhraseTableProps {
    phrases: Phrase[];
    onPhraseChange: (index: number, field: keyof Phrase, value: string) => void;
    onAddPhrase: () => void;
    onRemovePhrase: (index: number) => void;
}

const PhraseTable: React.FC<PhraseTableProps> = ({ phrases, onPhraseChange, onAddPhrase, onRemovePhrase }) => {
    return (
        <TableContainer component={Box} sx={{ mt: 2 }}>
            <Table>
                <TableHead>
                    <TableRow>
                        <TableCell sx={{ whiteSpace: 'nowrap', minWidth: 150 }}>フレーズ（必須）</TableCell>
                        <TableCell sx={{ whiteSpace: 'nowrap', minWidth: 150 }}>みたいに聞こえる（任意）</TableCell>
                        <TableCell sx={{ whiteSpace: 'nowrap', minWidth: 150 }}>IPA（任意）</TableCell>
                        <TableCell sx={{ whiteSpace: 'nowrap', minWidth: 150 }}>表示フレーズ（任意）</TableCell>
                        <TableCell></TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {phrases.map((row, index) => (
                        <TableRow key={index}>
                            <TableCell>
                                <TextField
                                    value={row.phrase}
                                    onChange={(e) => onPhraseChange(index, 'phrase', e.target.value)}
                                    required
                                    fullWidth
                                />
                            </TableCell>
                            <TableCell>
                                <TextField
                                    value={row.soundsLike}
                                    onChange={(e) => onPhraseChange(index, 'soundsLike', e.target.value)}
                                    fullWidth
                                />
                            </TableCell>
                            <TableCell>
                                <TextField
                                    value={row.ipa}
                                    onChange={(e) => onPhraseChange(index, 'ipa', e.target.value)}
                                    fullWidth
                                />
                            </TableCell>
                            <TableCell>
                                <TextField
                                    value={row.displayAs}
                                    onChange={(e) => onPhraseChange(index, 'displayAs', e.target.value)}
                                    fullWidth
                                />
                            </TableCell>
                            <TableCell>
                                <IconButton onClick={() => onRemovePhrase(index)}>
                                    <DeleteIcon />
                                </IconButton>
                            </TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
            <Box display="flex" justifyContent="flex-end" p={2}>
                <Button variant="contained" color="primary" onClick={onAddPhrase}>
                    行を追加
                </Button>
            </Box>
        </TableContainer>
    );
};

const UploadCSV: React.FC<{ onUpload: (phrases: Phrase[]) => void; fileName: string | null; setFileName: React.Dispatch<React.SetStateAction<string | null>> }> = ({
                                                                                                                                                                       onUpload,
                                                                                                                                                                       fileName,
                                                                                                                                                                       setFileName
                                                                                                                                                                   }) => {
    const [openSnackbar, setOpenSnackbar] = useState(false);
    const [errorMessage, setErrorMessage] = useState<string | null>(null);

    const handleFileChange = async (event: React.ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.[0];
        if (!file) return;

        setFileName(file.name); // ファイル名を親コンポーネントの状態として設定

        try {
            const reader = new FileReader();
            reader.onload = (e) => {
                const arrayBuffer = e.target?.result as ArrayBuffer;
                const uint8Array = new Uint8Array(arrayBuffer);
                const workbook = read(uint8Array, { type: 'array' });
                const sheetName = workbook.SheetNames[0];
                const worksheet = workbook.Sheets[sheetName];
                const data = utils.sheet_to_json<Phrase>(worksheet);
                onUpload(data); // データを親コンポーネントに渡す
                setOpenSnackbar(true); // 成功メッセージを表示
            };
            reader.readAsArrayBuffer(file); // readAsArrayBufferを使用
        } catch (error) {
            setErrorMessage('ファイルの読み込みに失敗しました。');
            console.error('Error reading file:', error);
        }
    };

    const handleRemoveFile = () => {
        setFileName(null);
        setErrorMessage(null);
        onUpload([]); // アップロードされたデータをリセット
    };

    return (
        <Box sx={{ mt: 2 }}>
            <Button
                variant="contained"
                component="label"
                startIcon={<UploadFileIcon />}
                sx={{ mb: 2, backgroundColor: 'green', color: 'white', '&:hover': { backgroundColor: 'darkgreen' } }}
            >
                CSVファイルを選択
                <input type="file" accept=".csv, .xlsx" hidden onChange={handleFileChange} />
            </Button>
            {fileName && ( // ファイル名がある場合に表示
                <Card sx={{ display: 'flex', alignItems: 'center', mt: 2, p: 1, maxWidth: 400 }}>
                    <CardContent sx={{ flexGrow: 1 }}>
                        <Typography variant="body2">アップロードしたファイル:</Typography>
                        <Typography variant="body1" fontWeight="bold">{fileName}</Typography>
                    </CardContent>
                    <IconButton onClick={handleRemoveFile} sx={{color: 'green'}}>
                        <CloseIcon />
                    </IconButton>
                </Card>
            )}
            <Snackbar
                open={openSnackbar}
                autoHideDuration={6000}
                onClose={() => setOpenSnackbar(false)}
                anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
            >
                <Alert onClose={() => setOpenSnackbar(false)} severity="success" sx={{ width: '100%' }}>
                    ファイルが正常にアップロードされました。
                </Alert>
            </Snackbar>
            {errorMessage && (
                <Snackbar
                    open={!!errorMessage}
                    autoHideDuration={6000}
                    onClose={() => setErrorMessage(null)}
                    anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
                >
                    <Alert onClose={() => setErrorMessage(null)} severity="error" sx={{ width: '100%' }}>
                        {errorMessage}
                    </Alert>
                </Snackbar>
            )}
        </Box>
    );
};

const DictionarySettings: React.FC = () => {
    const [tabIndex, setTabIndex] = useState(0); // タブのインデックス
    const [phrases, setPhrases] = useState<Phrase[]>([
        { phrase: '', soundsLike: '', ipa: '', displayAs: '' }
    ]);
    const [uploadedPhrases, setUploadedPhrases] = useState<Phrase[]>([]); // アップロードされたフレーズ用の新しい状態
    const [fileName, setFileName] = useState<string | null>(null); // アップロードされたファイル名を管理

    const handlePhraseChange = (index: number, field: keyof Phrase, value: string) => {
        const newPhrases = [...phrases];
        newPhrases[index][field] = value;
        setPhrases(newPhrases);
    };

    const handleAddPhrase = () => {
        setPhrases([...phrases, { phrase: '', soundsLike: '', ipa: '', displayAs: '' }]);
    };

    const handleRemovePhrase = (index: number) => {
        setPhrases(phrases.filter((_, i) => i !== index));
    };

    const handleUpload = (uploadedPhrases: Phrase[]) => {
        setUploadedPhrases(uploadedPhrases); // アップロードされたフレーズの状態のみ変更
    };

    return (
        <Paper elevation={3} sx={{ p: 4, mb: 4 }}>
            <Typography variant="h5" gutterBottom>
                辞書の設定
            </Typography>
            <Tabs value={tabIndex} onChange={(_, newValue) => setTabIndex(newValue)}>
                <Tab label="直接入力" />
                <Tab label="CSVアップロード" />
            </Tabs>
            {tabIndex === 0 && (
                <PhraseTable
                    phrases={phrases}
                    onPhraseChange={handlePhraseChange}
                    onAddPhrase={handleAddPhrase}
                    onRemovePhrase={handleRemovePhrase}
                />
            )}
            {tabIndex === 1 && (
                <UploadCSV
                    onUpload={handleUpload}
                    fileName={fileName}
                    setFileName={setFileName}
                />
            )}
        </Paper>
    );
};

export default DictionarySettings;
