import { CardMedia } from "@mui/material";
import { useTheme } from "@mui/material/styles";

const BusinessImage = () => {
  const theme = useTheme();
  return (
    <CardMedia
      component="img"
      sx={{
        height: theme.header.height,
        objectFit: "cover",
        objectPosition: "center",
      }}
      image={"/resources/images/business-image.png"}
      alt="Business Image"
    />
  );
};

export default BusinessImage;
