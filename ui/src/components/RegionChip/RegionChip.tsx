import LanguageIcon from '@mui/icons-material/Language';
import { Chip } from '@mui/material';

import UserRegion from '@/core/regions';

type RegionChipProps = {
  region: UserRegion;
};

function RegionChip({ region }: RegionChipProps) {
  return <Chip icon={<LanguageIcon />} label={region.getSettingName()} variant="outlined" />;
}

export default RegionChip;
