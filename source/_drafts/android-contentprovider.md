---
title: android_contentprovider
tags:
---
# Android ContentProvider简介
<!-- toc -->
## URI
 * '*' may be used as a wild card for any text
 * '#' may be used as a wild card for numbers
内容提供者是一种封装机制，而不是数据的访问机制。你需要一个实际的数据访问如SQLite或通过网络机制访问底层数据源。因此，内容提供者只是在应用程序间共享数据的一种抽象。对于内部数据，应用程序可以使用任何的数据存储/访问机制，一下是一些合适的方式：
 * Preferences: 可以存储键值对的数据。
 * Files: 文件存储
 * SQLite: 关系数据库，每个应用程序中都能创建一个私有的数据库。
 * Network: 通过网络检索或者存储数据。

## 如何查看数据库
可以使用adb命令在已连接的设备上打开shell
```zsh
$ adb shell
```

需要访问这个文件夹查看数据库列表
```zsh
ls /data/data
```

如果包含一个查找命令，你就可以查看所有的*.db文件。但是除了ls以外没有更好的方法了。最简便的方法是：
```zsh
ls  -R /data/data/*/databases
# 使用这个命令，你将看到Android列出所有的数据库（这里只列举常用，你的设备上可能还有很多）形式如下：
/data/data/com.ai.android.book.provider/databases:
bool.db
book.db
/data/data/com.android.alarmclock/databases:
alarms.db
/data/data/com.android.browser/databases:
browser.db
webviewCache.db
webview.db
/data/data/com.android.email/databases:
EmailProvider.db
EmailProviderBody.db
/data/data/com.android.globalsearch/databases:
shortcuts-log.db
/data/data/com.android.launcher/databases:
launcher.db
/data/data/com.android.providers.contacts/databases:
contacts2.db
/data/data/com.android.providers.downloads/databases:
downloads.db
/data/data/com.android.providers.media/databases:
external-ff32e0b.db
internal.db
/data/data/com.android.providers.settings/databases:
settings.db
/data/data/com.android.providers.telephony/databases:
telephony.db
mmssms.db
/data/data/com.pansi.msg/databases:
webview.db
webviewCache.db
pansi.db
```
这样就很方便的找到所有*.db数据
你可以调用sqlite3去查看这些数据库
```zsh
$ sqlite3 /data/data/com.android.providers.contacts/databases/contacts.db
```
## Sqlite快速入门
### SQL的指令格式 
SQL指令都是以分号（;）结尾的。如果遇到两个减号（--）则代表注解，sqlite3会略过去。
### 建立资料表 
假设我们要建一个名叫film的资料表，只要键入以下指令就可以了：
```sh
sqlite>create table film(title, length, year, starring);
```
这样我们就建立了一个名叫film的资料表，里面有name、length、year、starring四个字段。
这个create table指令的语法为：
```sh
sqlite>create table table_name(field1, field2, field3, ...);
```
table_name是资料表的名称，fieldx则是字段的名字。sqlite3与许多SQL数据库软件不同的是，它不在乎字段属于哪一种资料型态：sqlite3的字段可以储存任何东西：文字、数字、大量文字（blub），它会在适时自动转换。
### 查询资料 
键入以下命令查看数据库中所有的表：
```sh
sqlite>.tables
```
设置列头可见
```sh
sqlite>.headers on
```
我们首先简单介绍select的基本句型：
```sh
sqlite>select columns from table_name where expression;
```
最常见的用法，当然是倒出所有数据库的内容：
```sh
sqlite>select * from film;
```
如果资料太多了，我们或许会想限制笔数：
```sh
sqlite>select * from film limit 10;
```
或是照着电影年份来排列：
```sh
sqlite>select * from film order by year limit 10;
```
或是年份比较近的电影先列出来：
```sh
sqlite>select * from film order by year desc limit 10;
```
或是我们只想看电影名称跟年份：
```sh
select title, year from film order by year desc limit 10;
```
查所有茱蒂佛斯特演过的电影：
```sh
sqlite>select * from film where starring='Jodie Foster';
```
查所有演员名字开头叫茱蒂的电影('%' 符号便是 SQL 的万用字符）：
```sh
sqlite>select * from film where starring like 'Jodie%';
```
查所有演员名字以茱蒂开头、年份晚于1985年、年份晚的优先列出、最多十笔，只列出电影名称和年份：
```sh
sqlite>select title, year from film where starring like 'Jodie%' and year >= 1985 order by year desc limit 10;
```
有时候我们只想知道数据库一共有多少笔资料：
```sh
sqlite>select count(*) from film;
```
有时候我们只想知道1985年以后的电影有几部：
```sh
sqlite>select count(*) from film where year >= 1985;
```
### 在Android中查询数据库
### 使用URI读取数据
需要通过内容提供者提供的URI才能检索数据。因为这些URI定义都是唯一的，所以重要的是我们需要在调用之前查阅相关文档。Android中使用如下URI字符串提供了一些内容提供者的。
 
这些URI被定义在Android SDK的帮助类中：
 * !MediaStore.Images.Media.INTERNAL_CONTENT_URI
 * !MediaStore.Images.Media.EXTERNAL_CONTENT_URI
 * Contacts.People.CONTENT_URI
 
实际对应的值如下：
 * content://media/internal/images
 * content://media/external/images
 * content://contacts/people/
首先得获得URI，例如获得联系人People的URI
// content://contacts/people/
```java
Uri mContactsUri = Contacts.People.CONTENT_URI;
```
使用最好的方式检索数据
```java
Cursor cursor = getContentResolver().query(uri,null,//需要检索出那些列 projection
null, // 检索条件 selection
null,ContactsContract.Contacts.DISPLAY_NAME);//排序 order
```
请注意，projection是一个代表了列名的字符串数组。所以，除非你知道这些列名，否则你会发现很难创建projection。你可以再提供URI的同一个类中查看这些列名，由此，查看People类，你可以看到这些列名的定义：
 * DISPLAY_NAME
 * LAST_TIME_CONTACTED
 * NAME
 * NOTES
 * PHOTO_VERSION
 * SEND_TO_VOICE_MAIL
 * STARRED
 * TIMES_CONTACTED
例如：
```java
string[] projection = new string[] {
People._ID,
People.NAME,
People.NUMBER,
};
```
你可以在SDK文档的android.provider.Contacts.PerpleColumns类中查看到更多的列。你可以使用的如下[http://developer.android.com/reference/android/provider/Contacts.PeopleColumns.html Content URI]
selection就是sqlite语句中where之后的语句，比如只查询联系人名字以“电话”开头的项目(请注意空格（" like "）是必须的，否则会报错)：
```java
cursor = getContentResolver().query(uri,null,People.DISPLAY_NAME+" like "+"'电话%'", null,null);
```
而最后一项oder就是sqlite语句中"order by "之后的语句，比如我想已People.DISPLAY_NAME倒序排列，应如下：
```java
cursor = getContentResolver().query(uri,null,null, null,ContactsContract.Contacts.DISPLAY_NAME+" desc");
```
让我们重新回到游标，它包含了零条或者多条记录。列顺序和类型特定于内容提供者。然而，每一行都有一个默认的列_id为每一个的唯一标识。
 
# 使用游标 
这里有关于游标的一些特点：
* 游标是行数据的集合
* 在使用游标读取数据之前，你需要使用moveToFirst()方法。因为游标的起始位置在第一行数据之前。
* 需要知道列名
* 需要知道列的类型
* 所有的操作都居于列号，所以你必须把列名转换成列号
* 游标是任意游标（你可以向前，向后读取，也可以越过几行进行读取）
* 因为游标是任意游标，所以你必须通过行数量来访问
* 最后访问完后，必须关闭游标
Android提供了一组方法用于操作游标。下面的代码用于判断游标是否是一个空游标，如果不为空，则逐行读取数据。
一般取得cursor后，如下取值：
```java
if(null != cursor){
    try {
        while(cursor.moveToNext()){
        int nameColumnIndex = cursor.getColumnIndex(ContactsContract.Contacts.DISPLAY_NAME);
        String name = cursor.getString(nameColumnIndex);
        //对取到的值做操作
        }                 
    } finally {
        cursor.close();
    }
}
```
# 实现ContentProvider
我们已经讨论了如何查询内容提供者的数据，但是还没有讨论如何定义一个内容提供者。要定义一个内容提供者，你需要继承android.content.ContentProvider以及实现以下方法：
 * query
 * insert
 * update
 * getType
 
在实现这些方法之前，你需要学习一些知识。我们将以示例详细的说明实现一个内容提供者的步骤：
 1. 准备数据库，URI，列名等等，然后创建一个定义类，定义所需要的元数据。
 2. 继承抽象类ContentProvider
 3. 实现query,insert,update,delete,和getType方法
 4. 在manifest文件中注册内容提供者
# 准备数据库 
首先，我们需要创建一个包含书籍信息的数据库。这个数据库（book）只包含一个叫做 books的表，表中的列有：name,isbn和author。这些列名被定义在了metadata类中。在这个示例中定义了一个名为 BookProviderMetaData的类用来定义常量信息，该类得代码如下：
```java
public class BookProviderMetaData {
    public static final String AUTHORITY = "com.androidbook.provider.BookProvider";
    public static final String DATABASE_NAME = "book.db";
    public static final int DATABASE_VERSION = 1;
    public static final String BOOKS_TABLE_NAME = "books";
    private BookProviderMetaData() {}
    //inner class describing BookTable
    public static final class BookTableMetaData implements BaseColumns {
        private BookTableMetaData() {}
        public static final String TABLE_NAME = "books";
        //uri and MIME type definitions
        public static final Uri CONTENT_URI = Uri.parse("content://" + AUTHORITY + "/books");
        public static final String CONTENT_TYPE = "vnd.android.cursor.dir/vnd.androidbook.book";
        public static final String CONTENT_ITEM_TYPE = "vnd.android.cursor.item/vnd.androidbook.book";
        public static final String DEFAULT_SORT_ORDER = "modified DESC";
        //Additional Columns start here.
        //string type
        public static final String BOOK_NAME = "name";
        //string type
        public static final String BOOK_ISBN = "isbn";
        //string type
        public static final String BOOK_AUTHOR = "author";
        //Integer from System.currentTimeMillis()
        public static final String CREATED_DATE = "created";
        //Integer from System.currentTimeMillis()
        public static final String MODIFIED_DATE = "modified";
    }
}
```
在BookProviderMetaData类中定义了权限名为：com.androidbook.provider.!BookProvider。这个将作为在Android manifest文件中的注册字符串。这个字符串构成了URI的前一部分。
 
在这个类中还包含了一个名为BookTableMetaData的内部类。在这个内部类中定义了一个标识books所有数据的URI。这个URI如下所示：
content://com.androidbook.provider.BookProvider/books
 
这个URI值保存在常量BookProviderMetaData.!BookTableMetaData.CONTENT_URI中。
 
在类BookTableMetaData中还定义了books表的所有数据以及单行数据的MIME类型。实现的内容提供者将会根据传入的URI返回这些常量。
 
然后BookTableMetaData中定义了一组列名：name,isbn,author ,createde（创建时间）和modified（最后一次更新时间），值得注意的是，在定义这些元数据的时候，类型要与数据库字段的类型保持一致。
 
BookTableMetaData类继承自BaseColumns类，在BaseColumns类中提供了标准的_id字段，这个字段代表了每行数据的ID。有了这些元数据的定义，我们已经实现了内容提供者。
# 扩展ContentProvider 
为了实现BookProvider内容提供者得示例还需要继承ContentProvider类并且重写onCreate方法来创建数据库，然后实现 query,insert,update,delete和getType方法。本节涵盖数据库建立的步骤，而下面的部分将会对于 query,insert,update,delete和getType方法进行讲解。这些代码如下：
```java
public class BookProvider extends ContentProvider
{
//Logging helper tag. No significance to providers.
private static final String TAG = "BookProvider";
 
//Setup projection Map
//Projection maps are similar to "as" (column alias) construct
//in an sql statement where by you can rename the
//columns.
private static HashMap<String, String> sBooksProjectionMap;
static
{
sBooksProjectionMap = new HashMap<String, String>();
sBooksProjectionMap.put(BookTableMetaData._ID,
BookTableMetaData._ID);
//name, isbn, author
sBooksProjectionMap.put(BookTableMetaData.BOOK_NAME,
BookTableMetaData.BOOK_NAME);
sBooksProjectionMap.put(BookTableMetaData.BOOK_ISBN,
BookTableMetaData.BOOK_ISBN);
sBooksProjectionMap.put(BookTableMetaData.BOOK_AUTHOR,
BookTableMetaData.BOOK_AUTHOR);
//created date, modified date
sBooksProjectionMap.put(BookTableMetaData.CREATED_DATE,
BookTableMetaData.CREATED_DATE);
sBooksProjectionMap.put(BookTableMetaData.MODIFIED_DATE,
BookTableMetaData.MODIFIED_DATE);
}
 
//Setup URIs
//Provide a mechanism to identify
//all the incoming uri patterns.
private static final UriMatcher sUriMatcher;
private static final int INCOMING_BOOK_COLLECTION_URI_INDICATOR = 1;
private static final int INCOMING_SINGLE_BOOK_URI_INDICATOR = 2;
static {
sUriMatcher = new UriMatcher(UriMatcher.NO_MATCH);
sUriMatcher.addURI(BookProviderMetaData.AUTHORITY, "books",
INCOMING_BOOK_COLLECTION_URI_INDICATOR);
sUriMatcher.addURI(BookProviderMetaData.AUTHORITY, "books/#",
INCOMING_SINGLE_BOOK_URI_INDICATOR);
}
/**
* Setup/Create Database
* This class helps open, create, and upgrade the database file.
*/
private static class DatabaseHelper extends SQLiteOpenHelper {
DatabaseHelper(Context context) {
super(context,
BookProviderMetaData.DATABASE_NAME,
null,
BookProviderMetaData.DATABASE_VERSION);
}
@Override
public void onCreate(SQLiteDatabase db)
{
Log.d(TAG,"inner oncreate called");
db.execSQL("CREATE TABLE " + BookTableMetaData.TABLE_NAME + " ("
+ BookTableMetaData._ID + " INTEGER PRIMARY KEY,"
+ BookTableMetaData.BOOK_NAME + " TEXT,"
+ BookTableMetaData.BOOK_ISBN + " TEXT,"
+ BookTableMetaData.BOOK_AUTHOR + " TEXT,"
+ BookTableMetaData.CREATED_DATE + " INTEGER,"
+ BookTableMetaData.MODIFIED_DATE + " INTEGER"
+ ");");
}
@Override
public void onUpgrade(SQLiteDatabase db, int oldVersion, int newVersion)
{
Log.d(TAG,"inner onupgrade called");
Log.w(TAG, "Upgrading database from version "
+ oldVersion + " to "
+ newVersion + ", which will destroy all old data");
db.execSQL("DROP TABLE IF EXISTS " +
BookTableMetaData.TABLE_NAME);
onCreate(db);
}
}
private DatabaseHelper mOpenHelper;
//Component creation callback
 
@Override
public boolean onCreate()
{
Log.d(TAG,"main onCreate called");
mOpenHelper = new DatabaseHelper(getContext());
return true;
}
@Override
public Cursor query(Uri uri, String[] projection, String selection,
String[] selectionArgs, String sortOrder)
{
SQLiteQueryBuilder qb = new SQLiteQueryBuilder();
switch (sUriMatcher.match(uri)) {
case INCOMING_BOOK_COLLECTION_URI_INDICATOR:
qb.setTables(BookTableMetaData.TABLE_NAME);
qb.setProjectionMap(sBooksProjectionMap);
break;
case INCOMING_SINGLE_BOOK_URI_INDICATOR:
qb.setTables(BookTableMetaData.TABLE_NAME);
qb.setProjectionMap(sBooksProjectionMap);
qb.appendWhere(BookTableMetaData._ID + "="
+ uri.getPathSegments().get(1));
break;
default:
throw new IllegalArgumentException("Unknown URI " + uri);
}
// If no sort order is specified use the default
String orderBy;
if (TextUtils.isEmpty(sortOrder)) {
orderBy = BookTableMetaData.DEFAULT_SORT_ORDER;
} else {
orderBy = sortOrder;
}
// Get the database and run the query
SQLiteDatabase db = mOpenHelper.getReadableDatabase();
Cursor c = qb.query(db, projection, selection,
selectionArgs, null, null, orderBy);
//example of getting a count
int i = c.getCount();
// Tell the cursor what uri to watch,
// so it knows when its source data changes
c.setNotificationUri(getContext().getContentResolver(), uri);
return c;
}
@Override
public String getType(Uri uri)
{
switch (sUriMatcher.match(uri)) {
case INCOMING_BOOK_COLLECTION_URI_INDICATOR:
return BookTableMetaData.CONTENT_TYPE;
case INCOMING_SINGLE_BOOK_URI_INDICATOR:
return BookTableMetaData.CONTENT_ITEM_TYPE;
default:
throw new IllegalArgumentException("Unknown URI " + uri);
}
}
@Override
public Uri insert(Uri uri, ContentValues initialValues)
{
// Validate the requested uri
if (sUriMatcher.match(uri)!= INCOMING_BOOK_COLLECTION_URI_INDICATOR)
{
throw new IllegalArgumentException("Unknown URI " + uri);
}
ContentValues values;
if (initialValues != null) {
values = new ContentValues(initialValues);
} else {
values = new ContentValues();
}
Long now = Long.valueOf(System.currentTimeMillis());
// Make sure that the fields are all set
if (values.containsKey(BookTableMetaData.CREATED_DATE) == false)
{
values.put(BookTableMetaData.CREATED_DATE, now);
}
if (values.containsKey(BookTableMetaData.MODIFIED_DATE) == false)
{
values.put(BookTableMetaData.MODIFIED_DATE, now);
}
if (values.containsKey(BookTableMetaData.BOOK_NAME) == false)
{
throw new SQLException(
"Failed to insert row because Book Name is needed " + uri);
}
if (values.containsKey(BookTableMetaData.BOOK_ISBN) == false) {
values.put(BookTableMetaData.BOOK_ISBN, "Unknown ISBN");
}
if (values.containsKey(BookTableMetaData.BOOK_AUTHOR) == false) {
values.put(BookTableMetaData.BOOK_ISBN, "Unknown Author");
}
SQLiteDatabase db = mOpenHelper.getWritableDatabase();
long rowId = db.insert(BookTableMetaData.TABLE_NAME,
BookTableMetaData.BOOK_NAME, values);
if (rowId > 0) {
Uri insertedBookUri =
ContentUris.withAppendedId(
BookTableMetaData.CONTENT_URI, rowId);
getContext()
.getContentResolver()
.notifyChange(insertedBookUri, null);
return insertedBookUri;
}
throw new SQLException("Failed to insert row into " + uri);
}
@Override
public int delete(Uri uri, String where, String[] whereArgs)
{
SQLiteDatabase db = mOpenHelper.getWritableDatabase();
int count;
switch (sUriMatcher.match(uri)) {
case INCOMING_BOOK_COLLECTION_URI_INDICATOR:
count = db.delete(BookTableMetaData.TABLE_NAME,
where, whereArgs);
break;
case INCOMING_SINGLE_BOOK_URI_INDICATOR:
String rowId = uri.getPathSegments().get(1);
count = db.delete(BookTableMetaData.TABLE_NAME,
BookTableMetaData._ID + "=" + rowId
+ (!TextUtils.isEmpty(where) ? " AND (" + where + ')' : ""),
whereArgs);
break;
default:
throw new IllegalArgumentException("Unknown URI " + uri);
}
getContext().getContentResolver().notifyChange(uri, null);
return count;
}
@Override
public int update(Uri uri, ContentValues values,
String where, String[] whereArgs)
{
SQLiteDatabase db = mOpenHelper.getWritableDatabase();
int count;
switch (sUriMatcher.match(uri)) {
case INCOMING_BOOK_COLLECTION_URI_INDICATOR:
count = db.update(BookTableMetaData.TABLE_NAME,
values, where, whereArgs);
break;
case INCOMING_SINGLE_BOOK_URI_INDICATOR:
String rowId = uri.getPathSegments().get(1);
count = db.update(BookTableMetaData.TABLE_NAME,
values, BookTableMetaData._ID + "=" + rowId
+ (!TextUtils.isEmpty(where) ? " AND (" + where + ')' : ""),
whereArgs);
break;
default:
throw new IllegalArgumentException("Unknown URI " + uri);
}
getContext().getContentResolver().notifyChange(uri, null);
return count;
}
}
```
内容提供者的实现需要一个区分不同URI的机制。Android使用一个叫做UriMatcher的类来处理。所以我们需要使用所有包含的URI来构建此类对象。你可以在BookProvider类中创建projection map之后查看到代码示例。接下来我们将在“使用UriMatcher识别URI”节中对UriMatcher类进行讲解。
# 实现MIME类型 
BookProvider内容提供者必须实现getType()方法以返回给定URI的MIME类型。这个方法与一些其它内容提供者的方法一样——需要传入URI的重载方法。这个方法的首要任务是区分URI的类型。它是一个集合还是单行数据。
 
根据前面的内容，我们将使用UriMatcher去区分URI的类型。在BookTableMetaData类中已经定义了每类URI所需要返回的MIME类型。你可以查看在BookProvider类中查看这个方法的代码。
# 实现查询方法 
内容提供者的query方法根据URI和查询条件返回查询到得结果行的集合。
 
和其它方法一样，query方法使用UriMatcher去辨别URI类型。如果URI类型是单一项类型的（single-item type），这个方法通过如下方式获得book ID。
1.       使用getPathSegments()方法提取路径段
2.       第一个索引段即为book ID
query方法根据projections参数返回对应的列的数据。最后，query方法返回给调用者一个cursor对象。在这个过程中，query方法使用SQLiteQueryBuilder对象表示和执行查询。
# 注册内容提供者 
最后，你必须在Android.Manifest.xml文件中使用一下标签结构注册内容提供者：
```xml
<provider android:name=".BookProvider"
android:authorities="com.androidbook.provider.BookProvider"/>
```
