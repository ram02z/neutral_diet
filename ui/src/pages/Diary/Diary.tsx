import { FullSizeCenteredFlexBox } from '@/components/styled';
import useHeader from '@/store/header';

function Diary() {
  const [, headerActions] = useHeader();
  headerActions.change('Diary');

  return (
    <>
      <FullSizeCenteredFlexBox></FullSizeCenteredFlexBox>
    </>
  );
}

export default Diary;
