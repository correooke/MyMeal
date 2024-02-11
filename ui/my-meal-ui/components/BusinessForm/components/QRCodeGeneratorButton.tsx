import React from "react";
import Button from "@mui/material/Button";

const QRCodeGeneratorButton: React.FC = () => {
  const handleButtonClick = () => {
    // Add your code to generate QR code here
  };

  return (
    <Button variant="contained" onClick={handleButtonClick} color="secondary">
      Generar código QR
    </Button>
  );
};

export default QRCodeGeneratorButton;
