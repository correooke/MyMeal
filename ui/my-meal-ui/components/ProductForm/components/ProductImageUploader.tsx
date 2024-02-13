import React, { useState } from "react";
import { Box, CardMedia, Grid, Typography } from "@mui/material";
import { useTheme } from "@mui/material/styles";

const ProductImageUploader: React.FC = () => {
  const theme = useTheme();
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
          <CardMedia
            sx={{
              width: 120,
              height: 120,
              padding: 0,
              borderRadius: theme.border.radius,
              border: `1px solid ${theme.palette.primary.main}`,
            }}
            component="img"
            image={"/resources/images/photo-place-holder.png"}
            alt="Business Logo Image"
          />
        </label>
        {!selectedImage && (
          <Box maxWidth={120} textAlign={"center"}>
            <Typography variant="caption">Sube hasta dos imágenes</Typography>
          </Box>
        )}
        {selectedImage && (
          <Typography variant="body1">{selectedImage.name}</Typography>
        )}
      </Grid>
    </Grid>
  );
};

export default ProductImageUploader;
