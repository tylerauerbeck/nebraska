import '../../i18n/config.ts';

import { StyledEngineProvider, ThemeProvider } from '@mui/material/styles';
import { render } from '@testing-library/react';
import { describe, expect, it } from 'vitest';

import SearchInput from '../../components/common/ListSearch';
import themes from '../../lib/themes';

describe('List Search', () => {
  it('should render correct ListSearch', () => {
    const { asFragment } = render(
      <StyledEngineProvider injectFirst>
        <ThemeProvider theme={themes['light']}>
          <SearchInput value="test" />
        </ThemeProvider>
      </StyledEngineProvider>
    );
    expect(asFragment()).toMatchSnapshot();
  });
});
