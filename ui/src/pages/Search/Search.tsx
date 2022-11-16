import { FullSizeCenteredFlexBox } from '@/components/styled';
import useHeader from '@/store/header';

function Search() {
  const [, headerActions] = useHeader();
  headerActions.change('Search');

  return (
    <>
      <FullSizeCenteredFlexBox></FullSizeCenteredFlexBox>
    </>
  );
}

export default Search;
