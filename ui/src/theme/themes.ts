import { ThemeOptions } from '@mui/material/styles';
import { deepmerge } from '@mui/utils';

import { Themes } from './types';

const sharedTheme = {
  palette: {
    grey: {
      ['50']: '#fffbff',
      ['100']: '#f6eefa',
      ['200']: '#e7e0eb',
      ['300']: '#cbc4cf',
      ['400']: '#afa9b4',
      ['500']: '#948f99',
      ['600']: '#7a757f',
      ['700']: '#49454a',
      ['800']: '#322f37',
      ['900']: '#1d1a22',
    },
  },
  components: {
    MuiButtonBase: {
      defaultProps: {
        disableRipple: true,
      },
    },
    MuiDivider: {
      styleOverrides: {
        vertical: {
          marginRight: 10,
          marginLeft: 10,
        },
        // TODO: open issue for missing "horizontal" CSS rule
        // in Divider API - https://mui.com/material-ui/api/divider/#css
        middle: {
          marginTop: 10,
          marginBottom: 10,
          width: '80%',
        },
      },
    },
  },
} as ThemeOptions;

const themes: Record<Themes, ThemeOptions> = {
  light: deepmerge(sharedTheme, {
    palette: {
      mode: 'light',
      primary: {
        main: '#6f43c2',
        light: '#bb9aff',
        dark: '#5626a8',
      },
      error: {
        main: '#ba1a1a',
        light: '#ff897d',
        dark: '#93000a',
      },
      info: {
        main: '#7e525e',
        light: '#d39daa',
        dark: '#31101b',
      },
      success: {
        main: '#17a34c',
      },
      secondary: {
        main: '#635b70',
        light: '#b1a7bf',
        dark: '#1f182b',
      },
      background: {
        default: '#fffbff',
        paper: '#f6eefa',
      },
    },
    components: {
      MuiDrawer: {
        styleOverrides: {
          paper: {
            background: '#fffbff',
          },
        },
      },
    },
  }),

  dark: deepmerge(sharedTheme, {
    palette: {
      mode: 'dark',
      primary: {
        main: '#d3bbff',
        light: '#5626a7',
        dark: '#bb9aff',
      },
      error: {
        main: '#ffb4ab',
        light: '#690005',
        dark: '#ff897d',
      },
      info: {
        main: '#f0b7c5',
        light: '#643b4b',
        dark: '#ffd9e1',
      },
      success: {
        main: '#17a34c',
      },
      secondary: {
        main: '#cdc2db',
        light: '#4b4458',
        dark: '#e9def8',
      },
      background: {
        default: '#1d1b1e',
        paper: '#1d1a22',
      },
    },
    components: {
      MuiDrawer: {
        styleOverrides: {
          paper: {
            background: '#1d1b1e',
          },
        },
      },
      MuiAppBar: {
        styleOverrides: {
          colorPrimary: {
            backgroundColor: '#262329',
          },
        },
      },
      MuiBottomNavigation: {
        styleOverrides: {
          root: {
            backgroundColor: '#1d1b1e',
          },
        },
      },
    },
  }),
};

export default themes;
