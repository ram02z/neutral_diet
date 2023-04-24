import { useEffect, useState } from 'react';
import { isMobile } from 'react-device-detect';

import { useMediaQuery, useTheme } from '@mui/material';

/**
 * Determine the current layout of the application using the screen size and device type.
 */
function useLayout() {
  const [desktopLayout, setDesktopLayout] = useState(false);
  const theme = useTheme();
  const isMediumDevice = useMediaQuery(theme.breakpoints.down('md'));

  useEffect(() => {
    setDesktopLayout(!isMobile && !isMediumDevice);
  }, [isMediumDevice]);

  return { desktopLayout };
}

export default useLayout;
