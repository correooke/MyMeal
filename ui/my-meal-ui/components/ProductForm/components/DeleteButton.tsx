import React from "react";
import { Button } from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";

type DeleteButtonProps = {
  onClick: () => void;
};

const DeleteButton = ({ onClick }: DeleteButtonProps) => (
  <Button
    startIcon={<DeleteIcon />}
    onClick={() => onClick()}
    variant="contained"
    color="error"
  ></Button>
);

export default DeleteButton;
