import { Box, Typography } from "@mui/material";
import { Product } from "@/models/product";
import { useTheme } from "@mui/material/styles";

type ProductDetailsProps = {
  product: Product;
};

const ProductDetails = ({ product }: ProductDetailsProps) => {
  const theme = useTheme();
  return (
    <Box
      sx={{
        display: "flex",
        height: "100%",
        border: `1px solid ${theme.palette.primary.main}`,
      }}
    >
      <Typography variant="body1">Detalles</Typography>
    </Box>
  );
};

export default ProductDetails;
