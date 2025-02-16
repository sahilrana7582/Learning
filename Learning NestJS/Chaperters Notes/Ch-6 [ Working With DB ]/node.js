const generateUsers = (num) => {
    const users = [];
    for (let i = 1; i <= num; i++) {
      users.push({
        "firstName": `"FirstName${i}"`,
        "lastName": `"LastName${i}"`,
        "age": `"${(Math.random() * 100).toFixed(2)}"`,
        "nickName": `"NickName${i}"`,
        "password": `"password${i}"`
      });
    }
    console.log(JSON.stringify(users));
    return users;
  };
  
  const userArray = generateUsers(100);
  console.log(userArray);