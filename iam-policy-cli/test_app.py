import boto3

def download_user_avatar():
    # Connect to AWS S3
    s3 = boto3.client('s3')
    
    # Download the avatar image
    response = s3.get_object(
        Bucket='my-production-app-images',
        Key='avatars/user_123.png'
    )
    
    print("Download complete!")