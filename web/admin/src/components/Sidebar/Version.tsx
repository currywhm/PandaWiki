import { Box, Stack } from '@mui/material';
import packageJson from '../../../package.json';

const Version = () => {
  const curVersion = import.meta.env.VITE_APP_VERSION || packageJson.version;

  return (
    <>
      <Stack
        justifyContent={'center'}
        gap={0.5}
        sx={{
          borderTop: '1px solid',
          borderColor: 'divider',
          pt: 2,
          mt: 1,
          color: 'text.primary',
          fontSize: 12,
        }}
      >
        <Stack direction={'row'} gap={0.5}>
          <Box sx={{ width: 30, color: 'text.tertiary' }}>版本</Box>
          <Box sx={{ whiteSpace: 'nowrap' }}>{curVersion}</Box>
        </Stack>
      </Stack>
    </>
  );
};

export default Version;
