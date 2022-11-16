import useHeader from '@/store/header';
import { FullSizeCenteredFlexBox } from '@/components/styled';

function Home() {
  // FIXME: is this really the best way in React?
  const [,headerActions] = useHeader();
  headerActions.change("Home")

  return (
    <>
      <FullSizeCenteredFlexBox>
      </FullSizeCenteredFlexBox>
    </>
  );
}

export default Home;
