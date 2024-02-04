import React, { useState } from "react";
import { Button, Grid, Typography } from "@mui/material";

const ProductImageUploader: React.FC = () => {
  const [selectedImage, setSelectedImage] = useState<File | null>(null);

  const handleImageUpload = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0] ?? null;
    setSelectedImage(file);
  };

  return (
    <Grid container spacing={2}>
      <Grid item xs={12}>
        <input
          accept="image/*"
          id="product-image-upload"
          type="file"
          onChange={handleImageUpload}
          style={{ display: "none" }}
        />
        <label htmlFor="product-image-upload">
          <Button variant="contained" component="span">
            Selecciona la imagen
          </Button>
        </label>
        {!selectedImage && (
          <Typography variant="body1">Sube hasta dos imágenes</Typography>
        )}
        {selectedImage && (
          <Typography variant="body1">{selectedImage.name}</Typography>
        )}
      </Grid>
    </Grid>
  );
};

export default ProductImageUploader;
