import * as React from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import { Snackbar } from '@mui/material'
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import Box from '@mui/material/Box';


export default function GeneratePasswordDialog() {
  const [open, setOpen] = React.useState(false);
  const [length, setLength] = React.useState<number|string>("");
  const [requestedPassword, setRequestedPassword] = React.useState<boolean>(false);
  const [error, setError] = React.useState<boolean>(false);
  const [password, setPassword] = React.useState<string>("Loading...");

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
    setRequestedPassword(false);
    setError(false);
    setLength("");
  };
  const handleContinue = () => {
    if (length > 7 && length < 65) {
      // GET request for password
      setPassword("Super Secret Password")
      setRequestedPassword(true);
    } else {
      setError(true);
    }
  };

  return (
    <div style={{ display: 'inline' }}>
      <Button variant="outlined" onClick={handleClickOpen}>
        New Password
      </Button>
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>Password Generation</DialogTitle>
        {!requestedPassword && <DialogContent>
          <DialogContentText>
            Please choose the length for your password:
          </DialogContentText>
          <TextField
            autoFocus
            margin="dense"
            id="name"
            type='number'
            label="Password Length"
            fullWidth
            value={length}
            variant="standard"
            error={error}
            helperText={!error? " ":"Should be between 8 and 64 characters"}
            onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
              let value = parseInt(event.target.value)
              if (event.target.value == "" || value < 65) setLength(value)}}
          />
        </DialogContent>}
        {requestedPassword && <DialogContent>
          <DialogContentText>
            Your suggested password is:
          </DialogContentText>
          <Box display="flex">
            <TextField
              autoFocus
              margin="dense"
              id="name"
              label="Password"
              fullWidth
              value={password}
              variant="standard"
              disabled
            />
            <CopyToClipboardButton
            password={password}/>
          </Box>
        </DialogContent>}
        <DialogActions>
          <Button onClick={handleClose}>Close</Button>
          {!requestedPassword && <Button onClick={handleContinue}>Submit</Button>}
        </DialogActions>
      </Dialog>
    </div>
  );
}

type CopyProps = {
  password :string;
} 

const CopyToClipboardButton = ({password}: CopyProps) => {
    const [open, setOpen] = React.useState(false)
    const handleClick = () => {
      setOpen(true)
      navigator.clipboard.writeText(password)
    }
    
    return (
        <>
          <Button onClick={handleClick}><ContentCopyIcon/></Button>
          <Snackbar
            open={open}
            onClose={() => setOpen(false)}
            autoHideDuration={2000}
            message="Copied to clipboard"
          />
        </>
    )
}