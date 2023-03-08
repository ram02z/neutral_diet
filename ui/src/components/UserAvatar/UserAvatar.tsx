import { Avatar, AvatarProps } from '@mui/material';

function UserAvatar(props: AvatarProps & { name: string }) {
  let initials = null;
  if (props.name) {
    const nameWords = props.name.split(' ');
    const firstNameInitial = nameWords[0].charAt(0);
    const lastNameInitial = nameWords.length > 1 ? nameWords[nameWords.length - 1].charAt(0) : '';
    initials = firstNameInitial + lastNameInitial;
  }

  return <Avatar {...props}>{initials}</Avatar>;
}

export default UserAvatar;
