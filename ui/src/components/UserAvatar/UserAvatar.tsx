import { Avatar, AvatarProps, Tooltip } from '@mui/material';

/**
 * Avatar using name initials.
 */
function UserAvatar(props: AvatarProps & { name: string }) {
  let initials = null;
  if (props.name) {
    const nameWords = props.name.split(' ');
    const firstNameInitial = nameWords[0].charAt(0);
    const lastNameInitial = nameWords.length > 1 ? nameWords[nameWords.length - 1].charAt(0) : '';
    initials = firstNameInitial + lastNameInitial;
  }

  return (
    <Tooltip title={props.name}>
      <Avatar {...props}>{initials}</Avatar>
    </Tooltip>
  );
}

export default UserAvatar;
