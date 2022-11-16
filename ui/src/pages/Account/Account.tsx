import { FullSizeCenteredFlexBox } from '@/components/styled';
import useHeader from '@/store/header';

function Account() {
  const [, headerActions] = useHeader();
  headerActions.change('Account');
  return (
    <>
      <FullSizeCenteredFlexBox></FullSizeCenteredFlexBox>
    </>
  );
}

export default Account;
