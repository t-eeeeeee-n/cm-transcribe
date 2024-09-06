import React from 'react';
import {
    TextField,
    Select,
    MenuItem,
    FormControl,
    InputLabel,
    Typography,
    Paper,
    FormHelperText
} from '@mui/material';

interface VocabularyFormProps {
    vocabularyName: string;
    setVocabularyName: (name: string) => void;
    languageCode: string;
    setLanguageCode: (code: string) => void;
}

const VocabularyForm: React.FC<VocabularyFormProps> = ({
                                                           vocabularyName,
                                                           setVocabularyName,
                                                           languageCode,
                                                           setLanguageCode
                                                       }) => {
    return (
        <Paper elevation={3} sx={{ p: 4, mb: 4 }}>
            <Typography variant="h5" gutterBottom>
                基本情報
            </Typography>

            {/* ボキャブラリー名の入力フィールド */}
            <FormControl fullWidth variant="outlined" margin="normal">
                <TextField
                    id="vocabulary-name"
                    variant="outlined"
                    fullWidth
                    value={vocabularyName}
                    onChange={(e) => setVocabularyName(e.target.value)}
                    label="名前"
                />
                <FormHelperText>
                    語彙名の長さは最大 200 文字です。使用できる文字: a～z、A～Z、0～9、ピリオド (.)、ダッシュ (-)、およびアンダースコア (_)。
                </FormHelperText>
            </FormControl>

            {/* 言語の選択フィールド */}
            <FormControl fullWidth variant="outlined" margin="normal">
                <InputLabel htmlFor="language-select">言語</InputLabel>
                <Select
                    labelId="language-select-label"
                    id="language-select"
                    value={languageCode}
                    onChange={(e) => setLanguageCode(e.target.value as string)}
                    label="言語"
                    variant={"outlined"}
                >
                    <MenuItem value="en-US">英語、米国 (en-US)</MenuItem>
                    <MenuItem value="ja-JP">日本語 (ja-JP)</MenuItem>
                    <MenuItem value="fr-FR">フランス語 (fr-FR)</MenuItem>
                    <MenuItem value="es-ES">スペイン語 (es-ES)</MenuItem>
                </Select>
            </FormControl>
        </Paper>
    );
};

export default VocabularyForm;