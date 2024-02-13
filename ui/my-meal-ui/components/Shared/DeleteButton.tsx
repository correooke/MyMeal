import React from "react";
import { Button, Typography } from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";

type DeleteButtonProps = {
  onClick: () => void;
  isPrimaryColor: boolean;
};

const icon = () => <DeleteIcon></DeleteIcon>;

const DeleteButton = ({ onClick, isPrimaryColor }: DeleteButtonProps) => (
  <Button
    startIcon={icon()}
    onClick={() => onClick()}
    variant="text"
    color={isPrimaryColor ? "primary" : "secondary"}
  >
    Eliminar
  </Button>
);

export default DeleteButton;
