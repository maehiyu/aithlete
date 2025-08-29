import { useState } from 'react';
import { useCreateUser } from '../hooks/useParticipant';
import { TextField, Button, Box, Typography, MenuItem } from '@mui/material';

const sportsList = ['サッカー', 'バスケ', '野球', 'テニス', 'バレー']; // 必要に応じて拡張

export function UserInitialRegister() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [role, setRole] = useState('user');
  const [sports, setSports] = useState<string[]>([]);
  const { mutate, isPending, isSuccess, isError, error } = useCreateUser();

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    mutate({ name, email, role, sports });
  };

  return (
    <Box maxWidth={400} mx="auto" mt={8}>
      <Typography variant="h5" mb={2}>ユーザー初期登録</Typography>
      <form onSubmit={handleSubmit}>
        <TextField
          label="名前"
          value={name}
          onChange={e => setName(e.target.value)}
          fullWidth
          required
          margin="normal"
        />
        <TextField
          label="メールアドレス"
          value={email}
          onChange={e => setEmail(e.target.value)}
          fullWidth
          required
          margin="normal"
        />
        <TextField
          select
          label="役割"
          value={role}
          onChange={e => setRole(e.target.value)}
          fullWidth
          margin="normal"
        >
          <MenuItem value="user">ユーザー</MenuItem>
          <MenuItem value="coach">コーチ</MenuItem>
        </TextField>
        <TextField
          select
          label="スポーツ（複数選択可）"
          value={sports}
          onChange={e => setSports(typeof e.target.value === 'string' ? e.target.value.split(',') : e.target.value)}
          SelectProps={{ multiple: true }}
          fullWidth
          margin="normal"
        >
          {sportsList.map(sport => (
            <MenuItem key={sport} value={sport}>{sport}</MenuItem>
          ))}
        </TextField>
        <Button type="submit" variant="contained" color="primary" fullWidth disabled={isPending} sx={{ mt: 2 }}>
          登録
        </Button>
        {isSuccess && <Typography color="success.main" mt={2}>登録が完了しました！</Typography>}
        {isError && <Typography color="error.main" mt={2}>{(error as Error).message}</Typography>}
      </form>
    </Box>
  );
}
